package gqlclient

import (
	"bytes"
	"context"
)

func (u *UseCase) DownloadAssets(ctx context.Context) error {
	u.li.Log.AppInfo(ctx, "download assets...")

	var buf bytes.Buffer
	if err := u.adap.DownloadAssets(&buf); err != nil {
		return err
	}
	if err := u.li.Arv.UnArchive(&buf, u.config.Workdir); err != nil {
		return err
	}
	u.Log("success!")

	return nil
}

func (u *UseCase) UploadAssets(ctx context.Context) error {
	u.li.Log.AppInfo(ctx, "upload assets...")
	archive, err := u.li.Arv.Archive(".")
	if err != nil {
		return err
	}
	if err := u.adap.UploadAssets(archive); err != nil {
		return err
	}
	u.Log("success!")

	return nil
}
