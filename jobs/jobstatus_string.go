// Code generated by "stringer -type jobStatus"; DO NOT EDIT

package jobs

import "fmt"

const jobStatusname = "statusUnknownstatusHealthystatusUnhealthystatusMaintenance"

var jobStatusindex = [...]uint8{0, 13, 26, 41, 58}

func (i jobStatus) String() string {
	if i < 0 || i >= jobStatus(len(jobStatusindex)-1) {
		return fmt.Sprintf("jobStatus(%d)", i)
	}
	return jobStatusname[jobStatusindex[i]:jobStatusindex[i+1]]
}
