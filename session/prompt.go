package session

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/evilsocket/islazy/tui"
	"github.com/manifoldco/promptui"

	"github.com/muraenateam/muraena/core"
	"github.com/muraenateam/muraena/log"
)

const (
	InvalidCommand = "invalid option, enter help for assistance"
)

func Prompt(s *Session) {

	for {
		templates := &promptui.PromptTemplates{
			Prompt:  "{{ . | }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		}

		prompt := promptui.Prompt{
			Label:     ">",
			Templates: templates,
			Validate:  validate,
		}

		result, err := prompt.Run()
		if err == promptui.ErrInterrupt {
			exit()
		} else if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch strings.ToLower(result) {
		case "h", "help":
			help()
		case "exit":
			exit()
		case "victims", "credentials":
			s.showTracking(result)
		}

	}
}

func validate(input string) error {
	switch strings.ToLower(input) {
	case
		"", "h",
		"help",
		"exit",
		"victims", "credentials":
		return nil
	}

	return errors.New(InvalidCommand)
}

func help() {
	log.Info("**************************************************************************")
	log.Info("Options")
	log.Info(" help: %s", tui.Bold("Prints this help"))
	log.Info(" exit: %s", tui.Bold("Exit from "+core.Name))
	log.Info(" victims: %s", tui.Bold("Show active victims"))
	log.Info(" credentials: %s", tui.Bold("Show collected credentials"))
	log.Info("")
	log.Info("")
	log.Warning("NOTE: This feature is not fully implemented yet. ")
	log.Warning(" Follow evolutions on https://github.com/muraenateam/muraena/issues/5")
	log.Info("**************************************************************************")

}

func exit() {
	prompt := promptui.Prompt{
		Label:     "Do you want to exit",
		IsConfirm: true,
		Default:   "n",
	}
	answer, _ := prompt.Run()
	if strings.ToLower(answer) == "y" {
		os.Exit(0)
	}
}

func (s *Session) showTracking(what string) {

	m, err := s.Module("tracker")
	if err != nil {
		log.Error("%s", err)
		return
	}

	m.Prompt(what)
}
