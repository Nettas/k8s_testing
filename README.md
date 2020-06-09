# k8s_testing
Go app with development tools to learn k8s functionalit

lightweight go app for k8s

Skaffold commands:
initiatlize:
skaffold init

Configuration skaffold.yaml was written
You can now run [skaffold build] to build the artifacts
or [skaffold run] to build and deploy
or [skaffold dev] to enter development mode, with auto-redeploy

flags:
to expose service on local host with a port
example:
skaffold dev --port-forward 

to force a new init:
skaffold init --force
