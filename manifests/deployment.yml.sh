apiVersion: apps/v1
kind: Deployment
metadata:
  name: $TEAM_NAME
spec:
  replicas: 1
  selector:
    matchLabels:
      app: $TEAM_NAME
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
  minReadySeconds: 5
  template:
    metadata:
      labels:
        app: $TEAM_NAME
    spec:
      containers:
      - name: app
        image: gcr.io/$PROJECT_ID/$TEAM_NAME:$GITHUB_SHA
        ports:
        - containerPort: 8080
        envFrom:
        - configMapRef:
            name: $TEAM_NAME-cm
        resources:
          requests:
            cpu: 250m
            memory: 512Mi
          limits:
            cpu: 250m
            memory: 512Mi
