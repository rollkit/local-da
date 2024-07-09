# This Kurtosis package spins up a local DA

def run(plan):
    local_da_port_number = 7980
    local_da_port_spec = PortSpec(
        number=local_da_port_number,
        transport_protocol="TCP",
        application_protocol="http",
    )
    local_da_ports = {
        "jsonrpc": local_da_port_spec,
    }
    local_da = plan.add_service(
        name="local-da",
        config=ServiceConfig(
            image="ghcr.io/rollkit/local-da:v0.2.1",
            ports=local_da_ports,
            public_ports=local_da_ports,
        ),
    )
    
    # Set the local DA address to return for rollups
    local_da_address = "http://{0}:{1}".format(
        local_da.ip_address, local_da.ports["jsonrpc"].number
    )

    return local_da_address