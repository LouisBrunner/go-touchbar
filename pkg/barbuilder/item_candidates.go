package barbuilder

type Candidates struct {
	CommonProperties

	Candidates []string
	ForString  string
	ForRange   [2]int

	AllowsCollapsing                 bool
	StartsCollapsed                  bool
	Visible                          bool
	AllowsTextInputContextCandidates bool
}

func (me *Candidates) isAnItem() {}
