apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
    - deployment.yml
    - service.yml
configMapGenerator:
    - name: $TEAM_NAME-cm
      envs:
        - .env
