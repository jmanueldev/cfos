package scheduler

type DecisionEngine struct {
    optimizer *Optimizer
}

func (d *DecisionEngine) Schedule(
    workload Workload,
    nodes []Node,
) Node {

    bestScore := -1.0
    var best Node

    for _, n := range nodes {

        score := d.optimizer.Score(workload, n)

        if score > bestScore {
            bestScore = score
            best = n
        }
    }

    return best
}