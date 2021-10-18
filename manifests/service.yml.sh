apiVersion: v1
kind: Service
metadata:
  name: $TEAM_NAME-service
spec:
  type: NodePort
  ports:
    - port: 80
      targetPort: 8080
  selector:
    app: $TEAM_NAME
