# Contributing
We'd love your help making logger the best wrapper around multiple golang logger!

## Setup
Fork, then clone the repository:
```
mkdir -p <THIS_PROJECT_ROOT_DIRECTORY>
cd <THIS_PROJECT_ROOT_DIRECTORY>
git clone git@github.com:your_github_username/logger.git
cd logger
git remote add upstream https://github.com/amitrai48/logger.git
git fetch upstream
```

Enable go module if you clone the project inside your `$GOPATH` directory
```
export GO111MODULE=on;
```

## Making Changes
```
cd <THIS_PROJECT_ROOT_DIRECTORY>
git checkout master
git fetch upstream
git rebase upstream/master
git checkout -b new_feature
```
Make your changes, then ensure that make lint and make test still pass. If you're satisfied with your changes, push them to your fork.
git push origin `new_feature`
Then use the GitHub UI to open a pull request.

At this point, you're waiting on us to review your changes. We try to respond to issues and pull requests within a few business days, and we may suggest some improvements or alternatives. Once your changes are approved, one of the project maintainers will merge them.

We're much more likely to approve your changes if you:
Add tests for new functionality.
Write a good commit message.
Maintain backward compatibility.

## Some Recommended Contribution
1. Add more popular logger library wrapper 
