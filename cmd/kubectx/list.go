package main

import (
	"fmt"
	"io"

	"facette.io/natsort"
	"github.com/fatih/color"
	"github.com/pkg/errors"

	"github.com/ahmetb/kubectx/cmd/kubectx/kubeconfig"
)

// ListOp describes listing contexts.
type ListOp struct{}

func (_ ListOp) Run(stdout, _ io.Writer) error {
	kc := new(kubeconfig.Kubeconfig).WithLoader(defaultLoader)
	defer kc.Close()
	if err := kc.Parse(); err != nil {
		return errors.Wrap(err, "kubeconfig error")
	}

	ctxs := kc.ContextNames()
	natsort.Sort(ctxs)

	// TODO support KUBECTX_CURRENT_FGCOLOR
	// TODO support KUBECTX_CURRENT_BGCOLOR

	currentColor := color.New(color.FgGreen, color.Bold)
	if useColors() {
		currentColor.EnableColor()
	} else {
		currentColor.DisableColor()
	}

	cur := kc.GetCurrentContext()
	for _, c := range ctxs {
		s := c
		if c == cur {
			s = currentColor.Sprint(c)
		}
		fmt.Fprintf(stdout, "%s\n", s)
	}
	return nil
}
