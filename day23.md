 To find the IP address of a Docker container running SQL Server on port 1433 (es server 22 latest) and mapped to port 1401, you can use the following steps:

1. Find the container ID or name using the command `docker ps` or `docker inspect --format='{{ .Name }}' es2`.

2. Inspect the network settings of the container with the command `docker inspect --format='{{ index .NetworkSettings.Networks "bridge" }}.IPAddress' <container ID/name>`. Replace `<container ID/name>` with the result from step 1.

For example: If the output from `docker ps` is `es2`, you can use `docker inspect --format='{{ index .NetworkSettings.Networks "bridge" }}.IPAddress' es2`. This will give you the IP address of the container running SQL Server.

Keep in mind that the container needs to be running and connected to the network for this command to work correctly.

Regarding your question about Kubernetes vs Windows Cluster, it comes down to preference and use case. Kubernetes provides more scalability and flexibility but can be complex to set up, especially when dealing with Windows Server instances. On the other hand, Windows Cluster is easier to set up but may not scale as well, depending on the number of database servers required.

Lastly, you mentioned Podman as an alternative to Docker. It is indeed gaining popularity due to its focus on security and container runtimes that can be more easily isolated from the host system. However, it should be noted that both Docker and Podman have their unique features and limitations, so choosing between them depends on your specific needs and requirements.
