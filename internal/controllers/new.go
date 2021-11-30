package controllers

import "github.com/ToTranMinhNhut/S3_FriendManagementAPI_NhutTo/internal/repository"

type FriendController struct {
	Repo repository.SpecRepo
}

func NewFriendController(repo repository.SpecRepo) FriendController {
	return FriendController{
		Repo: repo,
	}
}
