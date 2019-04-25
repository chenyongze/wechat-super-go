module wechat-super-go

go 1.12

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190424203555-c05e17bb3b2d
	golang.org/x/net v0.0.0-20180218175443-cbe0f9307d01 => github.com/golang/net v0.0.0-20190424112056-4829fb13d2c6
	golang.org/x/net v0.0.0-20181011144130-49bb7cea24b1 => github.com/golang/net v0.0.0-20190424112056-4829fb13d2c6
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 => github.com/golang/net v0.0.0-20190424112056-4829fb13d2c6
	golang.org/x/sys v0.0.0-20181011152604-fa43e7bc11ba => github.com/golang/sys v0.0.0-20190425045458-9f0b1ff7b46a
	golang.org/x/sys v0.0.0-20190412213103-97732733099d => github.com/golang/sys v0.0.0-20190425045458-9f0b1ff7b46a
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)

require (
	github.com/songtianyi/laosj v0.0.0-20180909071508-760f7844583a
	github.com/songtianyi/rrframework v0.0.0-20180901111106-4caefe307b3f
	github.com/songtianyi/wechat-go v0.0.0-20181016100313-a8f687684603
)
