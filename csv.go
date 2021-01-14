package tsv

type CSV struct {
	*parser
}

func OpenCSV(file string) (*CSV, error) {
	var p = newParser(",")

	var err = p.loadFile(file)
	if err != nil {
		return nil, err
	}

	var t = &CSV{}
	t.parser = p

	return t, nil
}
