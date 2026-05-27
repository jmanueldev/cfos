from neo4j import GraphDatabase

class Graph:

    def __init__(self, uri, user, password):

        self.driver = GraphDatabase.driver(
            uri,
            auth=(user, password)
        )

    def create_node(self, node_id):

        with self.driver.session() as session:

            session.run(
                "MERGE (n:Node {id:$id})",
                id=node_id,
            )