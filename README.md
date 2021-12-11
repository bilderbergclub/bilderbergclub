# bilderbergclub

[Bilderberg](https://en.wikipedia.org/wiki/Bilderberg_meeting) club for gophers.

Only participant will be allowed to pass.

```go
package yourpkg

func WhateverYouCreate() {
	// check that you are a participant, panic if not.
	bilderbergclub.MustApproveThisMember()

	println("Welcome back, the mighty of this world.")

	// Do whatever you can in this world.
}
```

ee790d495ea1b189e644c5508a0c3d261ef64e5223c01f617bfd1889fe420a19efbfc6d5c11b159314ee88c7cacc22c0914e80775483cbe66a95d065d84fb7fe
