[![Build Status](https://travis-ci.org/trungng92/goolang.svg?branch=master)](https://travis-ci.org/trungng92/goolang)

I created this just to test using git's pre-commit hook to help make better commits.

Basically, whenever you do a commit, git will run tests, check compilations, or run anything you want.
This way, you can guarantee that your commit at least runs. This is useful because it encourages your commits to be singular ideas and allows you to take better advantage of `git bisect` to find bugs and regressions.

In addition, having pre-commit hooks empowers you to "test partial commits", i.e. when you have a list of changes, but you want to make multiple commits out of the work tree.

The `stash` git docs has good info on this <https://git-scm.com/docs/git-stash>.

>You can use `git stash save --keep-index` when you want to make two or more commits out of the changes in the work tree, and you want to test each change before committing:
>```
>	# ... hack hack hack ...
>	$ git add --patch foo            # add just first part to the index
>	$ git stash save --keep-index    # save all other changes to the stash
>	$ edit/build/test first part
>	$ git commit -m 'First part'     # commit fully tested change
>	$ git stash pop                  # prepare to work on all other changes
>	# ... repeat above five steps until one commit remains ...
>	$ edit/build/test remaining parts
>	$ git commit foo -m 'Remaining parts'
>```

