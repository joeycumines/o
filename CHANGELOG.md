# [v1.1.0] - Unreleased

## Added

* This CHANGELOG.md file by [antifuchs].
* A pull request template. Thanks to [wireguy] for the debugging help
  in [#13](https://github.com/antifuchs/o/issues/13).
* API documentation around some unsafe API interactions. Thanks,
  [antifuchs]!
* This repo now is a go module and lists all the things needed to
  build and test it.

# [v1.0.0] - 2019-03-19

## Added

* This is the first (semver) release of the ring-buffer accountancy
  package `o`. It provides a way for users to implement their own
  ring-buffers without forcing them to type-assert between types and
  `interface{}`.
* This release comes with an example (but serious) implementation of a
  `ReadWriter` that is backed by a ring buffer, in `ringio`.
* Most of the code in this repository by [antifuchs] & inspired by
  [a blog post](https://www.snellman.net/blog/archive/2016-12-13-ring-buffers/)
  by [jsnell].

## Fixed

* An issue with `maskRing` where `Shift`ing more times than the ring
  had capacity would return invalid indexes. Thanks for the bug
  report, [jsnell]!

<!-- github short links to contributors' profiles: -->
[antifuchs]: https://github.com/antifuchs
[jsnell]: https://github.com/jsnell
[wireguy]: https://github.com/wireguy

<!-- release version number short links: -->
[v1.0.0]: https://github.com/antifuchs/o/releases/tag/v1.0.0