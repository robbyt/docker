package node

import (
	"fmt"

	"github.com/docker/docker/api/client"
	"github.com/docker/docker/cli"
	"github.com/docker/engine-api/types/swarm"
	"github.com/spf13/cobra"
)

func newPromoteCommand(dockerCli *client.DockerCli) *cobra.Command {
	return &cobra.Command{
		Use:   "promote NODE [NODE...]",
		Short: "Promote a node to a manager in the swarm",
		Args:  cli.RequiresMinArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runPromote(dockerCli, args)
		},
	}
}

func runPromote(dockerCli *client.DockerCli, nodes []string) error {
	promote := func(node *swarm.Node) {
		node.Spec.Role = swarm.NodeRoleManager
	}
	success := func(nodeID string) {
		fmt.Fprintf(dockerCli.Out(), "Node %s promoted to a manager in the swarm.\n", nodeID)
	}
	return updateNodes(dockerCli, nodes, promote, success)
}
