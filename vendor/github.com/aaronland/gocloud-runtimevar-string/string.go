package runtimevar

import (
	"context"
	"errors"
	"github.com/aaronland/go-aws-session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	gc_runtimevar "gocloud.dev/runtimevar"
	"net/url"
	"path/filepath"
	"strings"
)

func OpenString(ctx context.Context, url_str string) (string, error) {

	if url_str == "" {
		return "", errors.New("Invalid URL string")
	}

	parsed, err := url.Parse(url_str)

	if err != nil {
		return "", err
	}

	query := parsed.Query()

	switch strings.ToUpper(parsed.Scheme) {
	case "AWSSECRETSMANAGER":

		aws_region := query.Get("region")
		aws_creds := query.Get("credentials")

		if aws_region == "" {
			return "", errors.New("Missing parameter: region")
		}

		if aws_creds == "" {
			return "", errors.New("Missing parameter: credentials")
		}

		aws_session, err := session.NewSessionWithCredentials(aws_creds, aws_region)

		if err != nil {
			return "", err
		}

		manager := secretsmanager.New(aws_session)

		secret_id := filepath.Join(parsed.Host, parsed.Path)

		result, err := manager.GetSecretValue(&secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secret_id),
		})

		if err != nil {
			return "", err
		}

		return *result.SecretString, nil

	default:

		decoder := query.Get("decoder")

		if decoder != "" && decoder != "string" {
			return "", errors.New("Invalid decoder")
		} else {
			query.Set("decoder", "string")

			parsed.RawQuery = query.Encode()
			url_str = parsed.String()
		}

		v, err := gc_runtimevar.OpenVariable(ctx, url_str)

		if err != nil {
			return "", err
		}

		defer v.Close()

		/*

			Latest is intended to be called per request, with the request context.
			It returns the latest good Snapshot of the variable value, blocking if
			no good value has ever been received. If ctx is Done, it returns the
			latest error indicating why no good value is available (not the ctx.Err()).
			You can pass an already-Done ctx to make Latest not block.

			https://godoc.org/gocloud.dev/runtimevar#Variable.Latest
		*/

		snapshot, err := v.Latest(ctx)

		if err != nil {
			return "", err
		}

		return snapshot.Value.(string), nil
	}

}
