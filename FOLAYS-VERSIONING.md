# VERSIONING

* Tags are immutable (once tagged, they should always point to the same git hash). That said, I will only enforce it **IF** `dear-imgui` **AND** `giu`  themselves enforce it. Otherwise, I will have no choice.
* Branches are mutable

### Quick fact

If you want to benefit of my work, for the time being, you will have to **MANUALLY** select my latest `git tag` according to the scheme below.

I will probably not have a `master branch` following both `dear-imgui` and `giu` because it could quickly pollute the repo (with all the `git subtree pull` + merge from both repos)

Instead, I will try to emit `semver`-versionned `git tag` based upon a `git subtree pull` each time done ontop of the previous release.

### Branches: (real, and examples)

My **`master`** branch will probably be synced to the latest `folays.imgui-v1.XX` (latest stable `giu` + ahead with latest stable `dear-imgui` if possible).


* `folays` (a `git fork` of `github.com/AllenDang/giu`), supposed to be kept in-sync with `giu`'s **`master`**
* `folays.imgui-v1.80` `giu`'s master + `imgui v1.80`
* `folays.imgui-v1.80.feature` : includes local feature (from `folays repository`)
* `folays.imgui-v1.81-WIP` which should be **ahead** of `giu 0.5.2 + dear-imgui 1.80` with updates from `dear-imgui 1.81`
* `folays.imgui-v1.81-WIP.feature` : includes local feature (from `folays` repository)
* `folays.imgui-v1.81` : future `giu`'s master + `imgui v1.81`
* `folays.imgui-v1.81.feature` : future `giu`'s master + `imgui v1.81` + local feature

### Tags:

* `v0.5.2-0.folays.imgui-v1.80.0.1-RELEASE` : **giu** `v0.5.2` with `imgui-v1.80`
* `v0.5.2-0.folays.imgui-v1.81.0.0-WIP.56f7bdae` :  **giu** `v0.5.2` with latest `imgui v.18-WIP` hash `56f7bdae`

I will try to make (if `git` allow me to do it):

* all `-RELEASE` be a `git subtree pull` of the previous `-RELEASE`.
* all '-WIP' be a `git subtree pull` of the nearest previous `-RELEASE`

### Git identifier deletion:

Git identifier that could be **DELETED** anytime :

* As `git subtree` use some sort of `squash`, and to not pollute the repo, branches could be deleted anytime. Please only use `tag` for stable `go.mod` / `go.sum`.
* The oldest `-WIP` git **`tags`** could also be deleted if its relevant

### Tree example:

```
v0.5.2-0.folays.imgui-v1.80.0.1-RELEASE
--------
^ imgui

v0.5.2-0.folays.imgui-v1.81.0.0-WIP.56f7bdae.feature
         ------                     --------
         ^ folays                   ^ WIP dear-imgui git hash

v0.5.2-0.folays.imgui-v1.81.0.1-RELEASE
                ˆˆˆˆˆˆˆˆˆˆˆˆˆ
                ^ dear-imgui x.y.z
                
v0.5.2-0.folays.imgui-v1.81.0.1-RELEASE.feature
                              ^ˆˆˆˆˆˆˆˆ Go semver : either 0-WIP or 1-RELEASE

v0.5.2-0.folays.imgui-v1.81.1.1-RELEASE.feature
                            -
                            ^ dear-imgui x.y.Z bugfix release (could be .1 or .b)
                                        -------            
                                        ^ local feature (from folays repository)
```

For the astute reader, special care have been observed to Go's `semver versioning` to have `dear-imgui`'s WIP have a lower precedence that final releases :

* Tags (`v*`) : special care
* Branches (`folays.*`) : special care not needed since they should not be used by `go get`

All tags are `semver` "pre-release" and as such, should not be selected by `go get`.

## Differences with github.com/AllenDang/giu

This repo is a git fork of github.com/AllenDang/giu

It aims to :

* use `git subtree` to sync from `github.com/ocornut/imgui` instead of hard-copied files (+ patches).

### BRANCH and TAGS `folays`

All branch and tags having the `semver` **`folays`** identifier have:

* the hard-copied `C++/header` files replaced by a `git subtree` instead
* but still have the `github.com/AllenDang/giu` patches applied

Semver identifier : [https://semver.org/#spec-item-9](https://semver.org/#spec-item-9)

# Components

* `github.com/ocornut/imgui` which is the original `dear-imgui` at the heart of the GUI library, code in C++
* `github.com/inkyblackness/imgui-go` which serves as a `cgo` bridge
* `github.com/AllenDang/giu` which propose an API more `Go idiomatic`

And

* a patch to `github.com/AllenDang/giu` (`power saving mode`)
* occasionally some patches specific to **my folays** fork repository (github.com/folays/giu)

Those latter patches would be mentionned either here, or in the branch/tag name (as documented in the `VERSIONING` section above)

# How it works

### `#include` usage (similar to a `symlink`)

1. `git subtree` feature impose a specific subdirectory (`/dear-imgui`).
2. `cgo` dictate that only same-directory (`Go package`) `.cpp/.h` are compiled.

Thus, to reconcile the above, `dear-imgui/*.{cpp,h}` files are `#include`'d in lieu of original `giu`'s `imgui/*.{cpp,h}` files.

`Symlink` do not work because `go download` ignore them. Thus, I `#include "../dear-imgui/"` them instead.

# Caveat

### `imgui-go`

`github.com/AllenDang/giu` seems to have substantially derived from `github.com/inkyblackness/imgui-go`
