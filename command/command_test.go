package command

import (
	"testing"

	"github.com/smallstep/cli/crypto/fingerprint"
)

func TestGetFingerprintEncoding(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name    string
		args    args
		want    fingerprint.Encoding
		wantErr bool
	}{
		{
			"hex",
			args{
				"HEX",
			},
			fingerprint.HexFingerprint,
			false,
		},
		{
			"base64",
			args{
				"base64",
			},
			fingerprint.Base64StdFingerprint,
			false,
		},
		{
			"base64url",
			args{
				"base64Url",
			},
			fingerprint.Base64URLFingerprint,
			false,
		},
		{
			"base64-url",
			args{
				"base64-URL",
			},
			fingerprint.Base64URLFingerprint,
			false,
		},
		{
			"base64raw",
			args{
				"base64raw",
			},
			fingerprint.Base64RawStdFingerprint,
			false,
		},
		{
			"base64-raw",
			args{
				"base64-raw",
			},
			fingerprint.Base64RawStdFingerprint,
			false,
		},
		{
			"base64urlraw",
			args{
				"base64urlraw",
			},
			fingerprint.Base64RawURLFingerprint,
			false,
		},
		{
			"base64url-raw",
			args{
				"base64url-raw",
			},
			fingerprint.Base64RawURLFingerprint,
			false,
		},
		{
			"base64-url-raw",
			args{
				"base64-url-raw",
			},
			fingerprint.Base64RawURLFingerprint,
			false,
		},
		{
			"emoji",
			args{
				"emoji",
			},
			fingerprint.EmojiFingerprint,
			false,
		},
		{
			"emojisum",
			args{
				"emojisum",
			},
			fingerprint.EmojiFingerprint,
			false,
		},
		{
			"unknown",
			args{
				"unknown",
			},
			fingerprint.HexFingerprint,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFingerprintEncoding(tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFingerprintEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFingerprintEncoding() got = %v, want %v", got, tt.want)
			}
		})
	}
}
