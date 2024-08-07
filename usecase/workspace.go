package usecase

import (
	"todo-app/domain/entity"
	"todo-app/domain/service"
)

type WorkspaceUseCase interface {
	GetAll() (workspaces []entity.Workspace, err error)
	GetByUserID(userID int) (workspaces []entity.Workspace, err error)
	Create(workspace entity.Workspace) (entity.Workspace, error)
}

type WorkspaceUseCaseImpl struct {
	workspace service.WorkspaceService
}

func NewWorkspaceUseCase(workspace service.WorkspaceService) *WorkspaceUseCaseImpl {
	return &WorkspaceUseCaseImpl{
		workspace,
	}
}

func (uc *WorkspaceUseCaseImpl) GetAll() ([]entity.Workspace, error) {
	workspaces, err := uc.workspace.GetAll()
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}

func (uc *WorkspaceUseCaseImpl) GetByUserID(userID int) ([]entity.Workspace, error) {
	workspaces, err := uc.workspace.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}

func (uc *WorkspaceUseCaseImpl) Create(workspace entity.Workspace) (entity.Workspace, error) {
	workspace, err := uc.workspace.Create(workspace)
	if err != nil {
		return entity.Workspace{}, err
	}

	return workspace, nil
}
