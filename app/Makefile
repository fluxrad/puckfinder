IMG=puckfinder-app

install:
	docker run -v "$(PWD)":/app $(IMG) npm install
	docker run -v "$(PWD)":/app $(IMG) npm run typings install
tsc:
	docker run -v "$(PWD)":/app $(IMG) npm run tsc
