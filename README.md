# Introduction
TODO: Write this

# Plan
- watch git repos for changes
  - Change defined as a branch changing sha, or a new tag. Possibly make this configurable
- when a change occurs, pull down the relevant sha
- Do a docker build on that sha
  - The arguments to pass to `docker build` should be configurable
- Push it up a registry v2 instance
  - ~~Do not tag by default. Create the image manifest, and push it up with the relevant layers, using the digest for the image. Need to push each layer first, then push the manifest~~
  - Have one tag that we keep pushing to for each build. Do not use `latest`.  It will return the digest for us to store.  That way, we do not need to reimplement `docker push`
- Keep a mapping between git sha and docker digest
  - It will have to be 1:many, since a git sha could result in different digests if `FROM` layer or installed packages change.
- Add the ability to add docker tags to certain digests automatically (branch name, git tag, git sha, `latest`)

# Future features
- Keep track of digest used for `FROM` layer. Allow users to recreate builds by using the same digest again, or using a newer digest
- Allow users to rebuild specific shas, to get different versions of packages or dependent layers
- Create a GUI frontend of some sort.  Perhaps in react or angular?