#!/bin/bash
set -x

# get all the node's ec2 instance ids for the specified chain id
chain_node_instance_ids=$(aws ec2 describe-instances --filters "Name=tag:NemoChainId,Values=$CHAIN_ID" | jq -r '[.Reservations | .[] | .Instances | .[] | .InstanceId] | join(" ")')

for chain_node_instance_id in ${chain_node_instance_ids}
do
    autoscaling_group_state=$(aws autoscaling describe-auto-scaling-instances --instance-ids "$chain_node_instance_id" | jq -r '[.AutoScalingInstances | .[].LifecycleState] | join(" ")')
    # Possible states: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-lifecycle.html
    case "$autoscaling_group_state" in
    Standby)
        # exit standby to make the node elgible for reciving
        # requests from the target groups for the public load balancer
        autoscaling_group_name=$(aws autoscaling describe-auto-scaling-instances --instance-ids "$chain_node_instance_id" | jq -r '[.AutoScalingInstances | .[].AutoScalingGroupName] | join(" ")')

        aws autoscaling exit-standby \
            --instance-ids "$chain_node_instance_id" \
            --auto-scaling-group-name "$autoscaling_group_name"
        ;;
    *)
        echo "instance ($chain_node_instance_id) not in an elgible state ($autoscaling_group_state) for exiting standby, skipping"
        ;;
    esac
done
