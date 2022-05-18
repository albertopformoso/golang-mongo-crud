create-db:
	docker-compose -f db/docker-compose.yml up -d --build

remove-db:
	docker-compose -f db/docker-compose.yml down

halt-db:
	make remove-db
	docker volume prune --force
	make create-db