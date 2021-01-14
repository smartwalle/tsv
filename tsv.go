package tsv

type TSV struct {
	*parser
}

func OpenTSV(file string) (*TSV, error) {
	var p = newParser("\t")

	var err = p.loadFile(file)
	if err != nil {
		return nil, err
	}

	var t = &TSV{}
	t.parser = p

	return t, nil
}
