package main

import "os/exec"

type IntermediateErr struct {
	error
}

func runjob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := isGloballyExec(jobBinPath)
	if err != nil {
		return IntermediateErr{wrapError(err, "can not run ", id)}
	} else if isExecutable == false {
		return wrapError(
			nil,
			"can not ru elseif",
			id,
		)
	}
	return exec.Command(jobBinPath, "--id="+id).Run()
}
