package app

func Run() {
	for {
		wd := getWorkingDir()

		selected := chooseFolder()
		dir, file := askPathParts()

		absFilePath := buildPath(wd, selected, dir, file)

		handleOverwrite(absFilePath)
		create(absFilePath)

		println("Created")
	}
}
