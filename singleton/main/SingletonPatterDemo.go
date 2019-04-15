package main

import "designPatter/singleton"

func main() {
	singletObjectNonSafety := singleton.GetInstanceNonSafety()
	singletObjectNonSafety.Output()

	singletObjectSafety := singleton.GetInstanceSafety()
	singletObjectSafety.Output()

	singletObjectSysSafety := singleton.GetInstanceSysSafety()
	singletObjectSysSafety.Output()

	singletObjectSelectSysSafety := singleton.GetInstanceSelectSynSafety()
	singletObjectSelectSysSafety.Output()
}
