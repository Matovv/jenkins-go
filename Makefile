.PHONY: git-push

gitpush:
	git add .
	git commit -m "fixes and updates"
	make git-push

git-push:
	git push origin