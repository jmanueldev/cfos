package optimizer

func Score(
    latency float64,
    thermal float64,
    carbon float64,
    cost float64,
) float64 {

    return (
        0.35 * (1.0 / latency) +
        0.25 * (1.0 / thermal) +
        0.25 * (1.0 / carbon) +
        0.15 * (1.0 / cost)
    )
}