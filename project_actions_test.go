package main

import "testing"

func TestAddAction(t *testing.T) {
	var projectLogic = &ProjectLogicMock{}

	project := Project{
		Name:    "test",
		Token:   "ABC",
		Service: "molly-test",
	}

	projectLogic.On("GenerateRandomToken").Return("123")
	projectLogic.On("HashToken", "123").Return(project.Token, nil)
	projectLogic.On("CreateFilesFolder", project).Return(nil)
	projectLogic.On("CreateDeploymentScript", project).Return(nil)
	projectLogic.On("CreateRunScript", project).Return(nil)
	projectLogic.On("CreateService", project).Return(nil)
	projectLogic.On("Save", project).Return(nil)

	var projectActions = ProjectActions{projectLogic}

	projectActions.AddAction(project.Name)

	projectLogic.AssertExpectations(t)
}

func TestStartServiceAction(t *testing.T) {
	var projectLogic = &ProjectLogicMock{}
	project := Project{}
	projectLogic.On("GetByName", "test", &project).Return(nil)
	projectLogic.On("RestartService", project).Return(nil)
	var projectActions = ProjectActions{projectLogic}

	projectActions.StartServiceAction("test")

	projectLogic.AssertExpectations(t)
}
