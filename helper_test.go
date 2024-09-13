// Copyright 2014-2022 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aerospike

func ParseInfoErrorCode(response string) Error {
	return parseInfoErrorCode(response)
}

func (e *AerospikeError) Msg() string {
	return e.msg
}

func (clstr *Cluster) GetMasterNode(partition *Partition) (*Node, Error) {
	return partition.getMasterNode(clstr)
}

// implements GomegaStringer to avoid some of the pain points
// in formatting the code
func (nd *Node) GomegaString() string {
	return nd.String()
}

func (ptn *Partition) GetMasterNode(cluster *Cluster) (*Node, Error) {
	return ptn.getMasterNode(cluster)
}

func (ptn *Partition) GetMasterProlesNode(cluster *Cluster) (*Node, Error) {
	return ptn.getMasterProlesNode(cluster)
}

// fillMinCounts will fill the connection pool to the minimum required
// by the ClientPolicy.MinConnectionsPerNode
func (nd *Node) ConnsCount() int {
	return nd.connectionCount.Get()
}

// CloseConnections closes all the node connections
func (nd *Node) CloseConnections() {
	nd.closeConnections()
}
