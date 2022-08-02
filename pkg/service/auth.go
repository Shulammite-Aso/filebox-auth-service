package service

import (
	"context"
	"errors"

	"github.com/Shulammite-Aso/filebox-auth-service/pkg/db"
	"github.com/Shulammite-Aso/filebox-auth-service/pkg/models"
	"github.com/Shulammite-Aso/filebox-auth-service/pkg/proto"
	"github.com/Shulammite-Aso/filebox-auth-service/pkg/utils"
)

type Server struct {
	proto.UnimplementedAuthServiceServer
	H   db.Handler
	Jwt utils.JwtWrapper
}

func (s *Server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Username: req.Username}).First(&user); result.Error == nil {
		return &proto.RegisterResponse{
			Message: "",
		}, errors.New("username not available. Please choose a differnt username")
	}

	user.Email = req.Email
	user.Username = req.Username
	user.Password = utils.HashPassword(req.Password)

	s.H.DB.Create(&user)

	return &proto.RegisterResponse{
		Message: "User created",
	}, nil
}

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Username}).First(&user); result.Error != nil {
		return &proto.LoginResponse{
			Token: "",
		}, errors.New("could not find user")
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &proto.LoginResponse{
			Token: "",
		}, errors.New("incorrect password")
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &proto.LoginResponse{
		Token: token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &proto.ValidateResponse{
			Username: "",
		}, err
	}

	var user models.User

	if result := s.H.DB.Where(&models.User{Username: claims.Username}).First(&user); result.Error != nil {
		return &proto.ValidateResponse{
			Username: "",
		}, errors.New("user not found")
	}

	return &proto.ValidateResponse{
		Username: user.Username,
	}, nil
}
