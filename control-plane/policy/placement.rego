package cfos

default allow = false

allow {
    input.node.region == input.intent.region
    input.node.carbon < 20
}