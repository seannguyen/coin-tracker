publish-build-base-image:
	docker build -t seannguyen/coin-tracker-build-base -f DockerfileBuildBase . ; \
	docker push seannguyen/coin-tracker-build-base;

publish-image:
	docker build -t seannguyen/coin-tracker . ; \
 	docker push seannguyen/coin-tracker;

deploy:
	 for host in $${DEPLOY_HOSTS} ; do \
		 cat scripts/pull_image_and_restart_container.sh | ssh -o "StrictHostKeyChecking no" $${DEPLOY_USERNAME}@$${host} -t; \
	 done

