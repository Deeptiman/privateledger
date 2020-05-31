<h1>Private Ledger</h1>
<p><a href="https://www.hyperledger.org/projects/fabric"><img src="https://www.hyperledger.org/wp-content/uploads/2016/09/logo_hl_new.png" alt="N|Solid"></a></p>
<p>Private Ledger is a web application written in Go to demonstrate storing private data in  Hyperleder fabric multi-org environment. The blockchain network consists of four organization joined with a single channel. The ledger data created in private collection, so that the data only accessible to the related organization unless the collection data is shared to other organization. And also, this repo will be demonstrate sharing the private collection data among the organizations.</p>
<p> However, this explanation guide does not explain how Hyperledger Fabric works, so for the information, you can follow at <a href="https://www.hyperledger.org/projects/fabric">Hyperledger.</a> </p>

<p><b>Medium writeup : </b><a href="https://medium.com/@deeptiman/confidentiality-and-private-data-in-hyperledger-fabric-1279c8e2e57f">https://medium.com/@deeptiman/confidentiality-and-private-data-in-hyperledger-fabric-1279c8e2e57f</a>

<h4><a id="Installation_6"></a>Installation</h4>
<p>Private Ledger requires <a href="https://www.docker.com/">Docker</a> &amp; <a href="https://golang.org/">Go</a> to run.</p>
<h3><a id="Docker_10"></a>Docker</h3>
<pre><code class="language-sh">$ sudo apt install docker.io
$ sudo apt install docker-compose
</code></pre>
<h2><a id="Go_15"></a>Go</h2>
<h4><a id="Installation_16"></a>Installation</h4>
<pre><code class="language-sh">$ sudo apt-get update
$ sudo apt-get install golang-go
</code></pre>
<h4><a id="Set_your_Go_path_as_environmental_variable_21"></a>Set your Go path as environmental variable</h4>
<h6><a id="add_these_following_variable_into_the_profile_22"></a>add these following variable into the profile</h6>
<pre><code class="language-sh">$ <span class="hljs-built_in">export</span> GOPATH=<span class="hljs-variable">$HOME</span>/go
$ <span class="hljs-built_in">export</span> PATH=<span class="hljs-variable">$PATH</span>:/usr/<span class="hljs-built_in">local</span>/go/bin:<span class="hljs-variable">$GOPATH</span>/bin
</code></pre>
<h6><a id="then_27"></a>then</h6>
<pre><code class="language-sh">$ <span class="hljs-built_in">source</span> ~/.profile
$ go version
$ go version go1.<span class="hljs-number">11</span> linux/amd64
</code></pre>

<h4> Setup the Host </h4>

<ul>
<li>
   Sometimes, in some of the local machine not able to identify or link the hyperledger endpoints with <b>localhost or 127.0.0.1</b>. So, it's better to add the hyperledger endpoints  in <b>/etc/hosts</b> mapping with 127.0.0.1   
</li>
  
  <pre><code>
  
    	127.0.0.1	orderer.private.ledger.com
	127.0.0.1	ca.org1.private.ledger.com
	127.0.0.1	peer0.org1.private.ledger.com
	127.0.0.1	peer1.org1.private.ledger.com
	127.0.0.1	peer0.org2.private.ledger.com
	127.0.0.1	peer1.org2.private.ledger.com
	127.0.0.1	peer0.org3.private.ledger.com
	127.0.0.1	peer1.org3.private.ledger.com
	127.0.0.1	peer0.org4.private.ledger.com
	127.0.0.1	peer1.org4.private.ledger.com
  
  </code></pre>
    
</ul>

<h4>Setup the Config</h4>
<ul>
<li>
<p>In a multi-org environment, few endorsement policies need to meet between all the organization to commit a transaction otherwise anybody can submit a transaction by signing as any member of the group.</p>
</li>
<li>
<p>The policies is mentioned in <b>configtx.yaml</b></p>
<ul>
<li>Readers, Writers &amp; Admins all have a specific Rules with Signing type</li>
</ul>
</li>
<li>
<p>Ex : <b>"OR('Org1MSP.member')"</b> means any memeber of Org1 can sign a transaction and commit blocks to the blockchain. There are various endorsing principals like admin, client, peer or member. The principals can be used based on the network architecture or business logic requirements.</p>
</li>
</ul>
<ol start="2">
<li>
<p><b>configtxgen</b> tool will create the channel artifacts and four Anchor peers of the organizations.</p>
</li>
<li>
<p>Anchor peers of all organization will be useful to communicate with each other so that any organization peer perform a transaction then other peers of the organization will get notified.</p>
</li>
<li>
<p>All the CAs, Peers and CouchDB for all the Orgs need to be mention in the docker-compose.yaml. In this repo, you can find there are four docker-compose YAML files, which are just the extended files from the base docker-compose.yaml. I have done this for code readability and changing config will be simpler.</p>
</li>
<li>
<p>Now the script to generate the artifacts mentioned below</p>
</li>
</ol>
<p><b>config.sh</b></p>
<pre><code> 

./bin/cryptogen generate --config=./crypto-config.yaml

./bin/configtxgen -profile FourOrgsOrdererGenesis -outputBlock ./artifacts/orderer.genesis.block

./bin/configtxgen -profile FourOrgsChannel -outputCreateChannelTx ./artifacts/privateledger.channel.tx -channelID privateledger

./bin/configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./artifacts/Org1MSPanchors.tx -channelID privateledger -asOrg Org1MSP

./bin/configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./artifacts/Org2MSPanchors.tx -channelID privateledger -asOrg Org2MSP

./bin/configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./artifacts/Org3MSPanchors.tx -channelID privateledger -asOrg Org3MSP

./bin/configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./artifacts/Org4MSPanchors.tx -channelID privateledger -asOrg Org4MSP

</code></pre>

<ol start="6">
<li>
<p>Start the network</p>
<pre><code>docker-compose up --force-recreate -d 
</code></pre>
</li>
</ol>
<p>This command will create the docker images for the organizations. It's needed to stop the existing running network and start the network as new, so I have mentioned the network down/up in a Makefile. Please use the make to run all at one instance.</p>
<h4>Deploy the network</h4>
<p>There are sequence steps to follow to completely setup the multi-org network.</p>
<ul>
<li>
<p>Initialize SDKs for each Organziations.</p>
<p>This steps will be called as an initial entry point to perform other activities in the network. Every organization will have its own SDK and will be initialized with config.yaml. The initialization process will create CA clients, MSP clients, Resource Management clients, signing identity.</p>
</li>
</ul>
<h5>1. Create Channel</h5>
<ul>
<li>The channel will create by Orderer and will use all the signing identities of the organization admins.</li>
</ul>
<pre><code>    

    req := resmgmt.SaveChannelRequest{
        ChannelID: "multiorgledger", 
        ChannelConfigPath: os.os.Getenv("GOPATH")+"/multiorgledger.channel.tx", 
        SigningIdentities: []msp.SigningIdentity{Org1SignIdentity, Org2SignIdentity, Org3SignIdentity, Org4SignIdentity},
    }
    
    txID, err := s.Resmgmt.SaveChannel(
        req, resmgmt.WithOrdererEndpoint(Orderer.OrdererID))
    
    if err != nil || txID.TransactionID == "" {
        return errors.WithMessage(err, "failed to save anchor channel for - "+s.OrgName)
    }
    
</code></pre>
<ul>    
<li>Each organization admins will create Anchor peers for their organization using the anchor peer artifacts.</li>
</ul>
<pre><code>   
        
        req := resmgmt.SaveChannelRequest{
               ChannelID: "multiorgledger", 
               ChannelConfigPath: os.os.Getenv("GOPATH")+"/", //Org1MSPanchors.tx or Org2MSPanchors.tx or Org3MSPanchors.tx or Org4MSPanchors.tx
                SigningIdentities: []msp.SigningIdentity{Org1SignIdentity or Org2SignIdentity or Org3SignIdentity or Org4SignIdentity},
        }
    
        txID, err := s.Resmgmt.SaveChannel(
                req, resmgmt.WithOrdererEndpoint(Orderer.OrdererID))
    
        if err != nil || txID.TransactionID == "" {
                return errors.WithMessage(err, "failed to save anchor channel for - "+s.OrgName)
         }
	 
</code></pre>
<h5>2. Join Channel</h5>
<ul>
<li>
<p>After creating the channel, each anchor peer will join the channel and remember Orderer will only create the channel and he can't join the channel.</p>
 <p>s.Resmgmt - is the resource management client created during Org SDK initialization. So, each organization resource management client will execute the join channel action indivisually.</p>
<pre><code>

  if err := s.Resmgmt.JoinChannel(s.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts),                                               resmgmt.WithOrdererEndpoint(Orderer.OrdererID)); err != nil {
     return errors.WithMessage(err, "failed to make admin join channel")
  }
  
</code></pre>
</li>
<li>
<p>There are specific channel queries can be use to check if the peers of an organization has already join the channel or not.</p>
<p>for this query specific peer details needs to be pass, so it will check for the channel join status.</p>
<pre><code> 

        resp, err := orgResmgmt.QueryChannels(resmgmt.WithTargets(peer))
        if err != nil {
              fmt.Println("IsJoinedChannel : failed to Query &gt;&gt;&gt; "+err.Error())
              return false, err
        }
 
        for _, chInfo := range resp.Channels {
              fmt.Println("IsJoinedChannel : "+chInfo.ChannelId+" --- "+s.ChannelID)
               if chInfo.ChannelId == s.ChannelID {
                        return true, nil
                }
        }
	
</code></pre>
</li>
</ul>
<h5>3. Install Chaincode</h5>
<ul>
<li>
<p>Installing chaincode for the organization will use the same chaincode id &amp; version unless it requires to be different for an organization on specific circumstances.</p>
</li>
<li>
<p><b>InstallCCRequest</b> will be the same for all organizations but the request will be used by the individual resource management client of the organizations. So, the chaincode will be installed in all the organization separately.</p>
<pre><code>  

        req := resmgmt.InstallCCRequest{
                Name:    s.ChaincodeId,
                Path:    s.ChaincodePath,
                Version: s.ChainCodeVersion,
                Package: ccPkg,
        }

        _, err = s.Resmgmt.InstallCC(req,
                resmgmt.WithRetry(retry.DefaultResMgmtOpts))

        if err != nil {
           fmt.Println("failed to install chaincode : "+err.Error())
           return errors.WithMessage(err, "  failed to install chaincode")
        }
	
</code></pre>
</li>
</ul>
<h5>4. Instantiate Chaincode</h5>
<ul>
<li>
<p>Chaincode instantiation in a multi org environment will be done only once. It will be initiated by any organization peers and will specify certain endorsement policy considering all the organization member or peer or admin depending on the network configurations.</p>
</li>
<li>
<p>As this application specifically developed to store private data in the chaincode, so each organization will have its policy and collections.</p>
</li>
<li>
<p><b>For Example</b></p>
      Collection : collectionOrg1

      Policy : "OR ('Org1MSP.member')"
<p>so this specifies that only and only "org1" can use the collectionOrg1 data. The data is completely restricted to other organizations. As this application mainly intended to store private data per organization and later share the data among organizations with certain access query. So, each organization will have there own private collection and policies and it will be added to the chaincode instantiation request.</p>
</li>
</ul>
<p>This function will return the collection config object per organization, so later each collection config can be pass as an array to the chaincode instantiation.</p>
<pre><code>

	func newCollectionConfig(colName, policy string, reqPeerCount, maxPeerCount int32, blockToLive uint64) (*cb.CollectionConfig, error) {

	        p, err := cauthdsl.FromString(policy)

        	if err != nil {
	            fmt.Println("failed to create newCollectionConfig : "+err.Error())
	            return nil, err
	        }
	        cpc := &amp;cb.CollectionPolicyConfig{
	            Payload: &amp;cb.CollectionPolicyConfig_SignaturePolicy{
	                SignaturePolicy: p,
	            },
	        }
	        return &amp;cb.CollectionConfig{
	            Payload: &amp;cb.CollectionConfig_StaticCollectionConfig{
	                StaticCollectionConfig: &amp;cb.StaticCollectionConfig{
	                    Name:              colName,
	                    MemberOrgsPolicy:  cpc,
	                    RequiredPeerCount: reqPeerCount,
	                    MaximumPeerCount:  maxPeerCount,
	                    BlockToLive:       blockToLive,
	                },
	            },
	        }, nil
	}

	// The maximum no of block allocated to keep private data alive, after that the data will purge across the network 
	blockToLive = 1000 
  // Number of peer required to distribute the private data as a condition of the endorsement of the chaincode
	peerCount = 0
  // This is the maxium number of peers allocated to a collection, if an endorsing peer goes down, then other peers are available
  maximumPeerCount = 3
	
  collCfg1, _ := newCollectionConfig("collectionOrg1", "OR ('Org1MSP.member')", peerCount, maximumPeerCount, blockToLive)
	collCfg2, _ := newCollectionConfig("collectionOrg2", "OR ('Org2MSP.member')", peerCount, maximumPeerCount, blockToLive)
	collCfg3, _ := newCollectionConfig("collectionOrg3", "OR ('Org3MSP.member')", peerCount, maximumPeerCount, blockToLive)
	collCfg4, _ := newCollectionConfig("collectionOrg4", "OR ('Org4MSP.member')", peerCount, maximumPeerCount, blockToLive)

	cfg := []*cb.CollectionConfig{ collCfg1, collCfg2, collCfg3, collCfg4}


<p>this collection config specifies that each collection record is private to the organization and other organization will only get hash of the ledger data.</p>
<p>// Any one of organization resource management client can execute the instantiate query and their peers will be mention in the target.
The chaincode policy will be adding all the member of the four organization</p>
   

	policy = "OR ('Org1MSP.member','Org2MSP.member','Org3MSP.member','Org4MSP.member')"     
	
	ccPolicy, _ := cauthdsl.FromString(policy) // cauthdsl will convert the policy string to Policy object

	resp, err := s.Resmgmt.InstantiateCC(
	    s.ChannelID,
	    resmgmt.InstantiateCCRequest{
 	       
	        Name:       s.ChaincodeId,
	        Path:       s.ChaincodePath,
	        Version:    s.ChainCodeVersion,
	        Args:       [][]byte{[]byte("init")},
	        Policy:     ccPolicy,
	        CollConfig: cfg,

	},resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithTargets(orgPeers[0], orgPeers[1]))

</code></pre>
<h5>5. Upgrade Chaincode</h5>
<ul>
<li>
<p>As the chaincode only be instantiated once, so if any changes made in the chaincode, then it will be upgraded with a new version code and keeping the chaincode name same (important) in the network.</p>
<p>// Any one of organization resource management client can execute the upgrade query and their peers will be mention in the target.</p>
<p>// Policy can be remain the same unless it requires modification based on your chaincode buiseness requirements.</p>
<pre><code>

	req := resmgmt.UpgradeCCRequest{
	  Name: 	      s.ChaincodeId, 
	  Version: 	    s.ChainCodeVersion, 
	  Path: 	      s.ChaincodePath, 
	  Args:  	      [][]byte{[]byte("init")},
	  Policy: 	    ccPolicy,
	  CollConfig:   cfg,    
	}

	resp, err := s.Resmgmt.UpgradeCC(s.ChannelID, req, resmgmt.WithRetry(retry.DefaultResMgmtOpts),resmgmt.WithTargets(orgPeers[0], orgPeers[1]))

	if err != nil {
	  return errors.WithMessage(err, " &gt;&gt;&gt;&gt; failed to upgrade chaincode")
	}
	
</code></pre>
</li>
</ul>
<h5>6. Query Installed Chaincode</h5>
<ul>
<li>
<p>Chaincode installation can be checked by querying "QueryInstalledChaincodes", which requires the peer of the organization, so it will check whether the peer has chaincode installed or not.</p>
<pre><code>  

	  resp, err := resMgmt.QueryInstalledChaincodes(resmgmt.WithTargets(peer))

	  if err != nil {
 	     return false, errors.WithMessage(err, "  QueryInstalledChaincodes for peer [%s] failed : "+peer.URL())
	  }
	  found := false

	  for _, ccInfo := range resp.Chaincodes {
	      fmt.Println("   "+orgID+" found chaincode "+ccInfo.Name+" --- "+ccName+ " with version "+ ccInfo.Version+" -- "+ccVersion)
      		if ccInfo.Name == ccName &amp;&amp; ccInfo.Version == ccVersion {
	          found = true
	          break
 	     }
	  }

	  if !found {
      		fmt.Println("   "+orgID+" chaincode is not installed on peer "+ peer.URL())
	        installedOnAllPeers = false
	  }  
	  
</code></pre>
</li>
</ul>
<h5>7. Query Instantiate Chaincode</h5>
<ul>
<li>Chaincode instantiation can be checked by querying "QueryInstantiatedChaincodes", which requires the peer of the organization, so it will check whether the peer has chaincode instantiated or not.</li>
</ul>
<pre><code>    


		chaincodeQueryResponse, err := resMgmt.QueryInstantiatedChaincodes(channelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts),    resmgmt.WithTargets(peer))

		    if err != nil {
		        return false, errors.WithMessage(err, "  QueryInstantiatedChaincodes return error")
 		    }
		    fmt.Println("\n   Found instantiated chaincodes on peer "+peer.URL())

		    found := false

		    for _, chaincode := range chaincodeQueryResponse.Chaincodes {
		        fmt.Println("   Found instantiated chaincode Name: "+chaincode.Name+", Version: "+chaincode.Version+", Path: "+chaincode.Path+" on peer "+peer.URL())
		        if chaincode.Name == ccName &amp;&amp; chaincode.Version == ccVersion {
		            found = true
		            break
		        }
 		   }

		    if !found {
		        fmt.Println("  "+ccName+" chaincode is not instantiated on peer "+ peer.URL())
		        installedOnAllPeers = false
		    } 
		    
</code></pre>
<h5>8. Affiliate an Org</h5>
<ul>
<li>
<p>This is most important in case CA registration. In Hyperledger fabric by default "org1 &amp; org2" are affiliated as CA organization, so any client or peer wants to register or enroll into the network via CA can pass "org1 or org2" as an affiliated organization.</p>
</li>
<li>
<p>But in case of other organization like org3 &amp; org4, they need to be affiliated using following CA Client API.</p>
<pre><code>

		// to perform the query individual org ca client needs to be used  

		affl := strings.ToLower(org) + ".department1"

		_, err = caClient.AddAffiliation(&amp;caMsp.AffiliationRequest{

		      Name:   affl,
		      Force:  true,
		      CAName: caid,
		})

		if err != nil {
			return fmt.Errorf("Failed to add affiliation for CA '%s' : %v ", caid, err)
		}
		
</code></pre>
<p>we can also check, whether the organization has the affiliation from the CA client by using the following API.</p>
<pre><code>

		fRes, err := caClient.GetAffiliation(affl)

		if afRes != nil &amp;&amp; err != nil {

			  fmt.Println("Affiliation Exists")

			  AfInfo := afRes.AffiliationInfo
			  CAName := afRes.CAName

			  fmt.Println("AfInfo : " + AfInfo.Name)
			  fmt.Println("CAName : " + CAName)
		}
		
</code></pre>
</li>
</ul>

<h4>Dependency Issues</h4>
<ol>
   <li>
      Hyperledger fabric-sdk-go is still in development. If you do dep ensure for each <b>Gopkg.toml</b> in <b>PrivateLedger</b> and <b>Chaincode</b>, it will download the govendor folder for each module but it will have some compilation issues while building the project. I have corrected the error for both <b>PrivateLedger and Chaincode</b> folder.
   </li>
   <li>
   Please download the vendor folder and add it in your project repo.   
      
   PrivateLedger - https://www.dropbox.com/s/ry1jmw0y9xliose/vendor.zip?dl=0
   
   Chaincode - https://www.dropbox.com/s/31nnqflpqwaywoa/vendor.zip?dl=0
   </li>
   <li>
   <b>Add vendor folders at the location where Gopkg.toml file is located.</b>
   </li>
</ol>

<p>So, this concludes the essential multi-org network setup for a 4 organizations based network.</p>
<h4>Run the application</h4>
<ol>
<li>
<p>The application is developed consisting of two clients. ( REST and Web App)</p>
</li>
<li>
<p>You can use the <b>REST</b> client at the server running at - <a href="http://localhost:4000">http://localhost:4000</a></p>
</li>
<li>
<p>And for <b>Web App</b>, you can use server running at - <a href="http://localhost:6000">http://localhost:6000</a></p>
</li>
</ol>

<h2>License</h2>
<p>This project is licensed under the <a href="https://github.com/Deeptiman/privateledger/blob/master/LICENSE">MIT License</a></p>
