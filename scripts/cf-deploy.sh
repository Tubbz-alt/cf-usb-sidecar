#! /usr/bin/env bash

# This script deploys the given service on the default kubernetes context

set -o errexit -o nounset
if test -r "${1:-}/Chart.yaml" ; then
    # Explicit path to chart, use that and autodetect type
    chart_path="${1:-}"
    service="$(grep name: "${chart_path}/Chart.yaml" | sed 's@name: cf-usb-sidecar-@@')"
else
    # Service by type
    service="$(tr '[:upper:]' '[:lower:]' <<<"${1:-mysql}")"
    chart_path="$(dirname "${0}")/../csm-extensions/services/dev-${service}/output/helm"
    if ! test -r "${chart_path}/Chart.yaml" ; then
        printf "Failed to find helm chart at %b%s%b\n" "\033[0;1;31m" "${chart_path}" "\033[0m" >&2
        exit 1
    fi
fi
shift 1

host="${service^^}"
extra=""
case "${service}" in
    mysql)
        ;;
    postgres)
        host="POSTGRESQL"
        extra="--set env.SERVICE_POSTGRESQL_SSLMODE=disable"
        ;;
    *)
        printf "Unknown service %s\n" "${service}" >&2
        exit 1
        ;;
esac

deployment_name="$(helm list | grep 'DEPLOYED' | awk '$NF == "cf" { print $1 }' | tail -n1)"

get_value() {
    helm get values "${deployment_name}" | y2j | jq -r "${1}"
}

get_secret() {
    kubectl get secret -n cf secret -o jsonpath="{@${1}}" | base64 -d
}

helm list --all | \
    awk "\$NF == \"dev-${service}\" { print \$1 }" | \
    xargs --no-run-if-empty helm delete --purge

printf "Deploying %b%s%b\n" "\033[0;1;32m" "dev-${service}" "\033[0m"

helm install \
    "${chart_path}" \
    --name "dev-${service}" \
    --namespace "dev-${service}" \
    --wait \
    --timeout 300 \
    --set env.CF_ADMIN_PASSWORD="$(get_secret .data.cluster-admin-password)" \
    --set env.CF_ADMIN_USER=admin \
    --set env.CF_CA_CERT="$(get_secret .data.internal-ca-cert)" \
    --set env.CF_DOMAIN="$(get_value .env.DOMAIN)" \
    --set env.SERVICE_LOCATION="http://cf-usb-sidecar-${service}.dev-${service}.svc.cluster.local:8081" \
    --set env.UAA_CA_CERT="$(get_secret .data.uaa-ca-cert)" \
    --set env.SERVICE_${host}_HOST=AUTO \
    --set kube.registry.hostname="${DOCKER_REPOSITORY:-docker.io}" \
    --set kube.organization="${DOCKER_ORGANIZATION:-splatform}" \
    ${extra} "$@"
