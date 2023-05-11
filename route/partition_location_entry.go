package route

import "strconv"

type obPartLocationEntry struct {
	partLocations map[int64]*obPartitionLocation
}

func newObPartLocationEntry(partNum int) *obPartLocationEntry {
	entry := new(obPartLocationEntry)
	entry.partLocations = make(map[int64]*obPartitionLocation, partNum)
	return entry
}

func (e *obPartLocationEntry) String() string {
	var partitionLocationStr string
	var i = 0
	partitionLocationStr = partitionLocationStr + "{"
	for k, v := range e.partLocations {
		if i > 0 {
			partitionLocationStr += ", "
		}
		i++
		partitionLocationStr += "m[" + strconv.Itoa(int(k)) + "]="
		if v != nil {
			partitionLocationStr += v.String()
		} else {
			partitionLocationStr += "nil"
		}
	}
	partitionLocationStr += "}"
	return "obPartLocationEntry{" +
		"partLocations:" + partitionLocationStr +
		"}"
}