publish-base-image:
	docker build -t seannguyen/coin-tracker-build-base -f DockerfileBuildBase . && \
	docker push seannguyen/coin-tracker-build-base;

publish-image:
	docker build -t seannguyen/coin-tracker . && \
 	docker push seannguyen/coin-tracker;

deploy:
	 for host in $${DEPLOY_HOSTS} ; do \
		ssh -o "StrictHostKeyChecking no" $${DEPLOY_USERNAME}@$${host} -t "/var/personal-infra/coin-tracker/start_docker.sh"; \
	 done

