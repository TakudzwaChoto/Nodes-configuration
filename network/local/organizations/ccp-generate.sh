#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $7)
    local CP=$(one_line_pem $8)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${PEER}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${P1PORT}/$4/" \
        -e "s/\${P2PORT}/$5/" \
        -e "s/\${CAPORT}/$6/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $7)
    local CP=$(one_line_pem $8)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${PEER}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${P1PORT}/$4/" \
        -e "s/\${P2PORT}/$5/" \
        -e "s/\${CAPORT}/$6/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

## prepare connection profile for orgwater_quality_data , updating and retrievals,storage
ORG=Water_quality_data,UpdatingAndRetrievals,Storage
PEER=water_quality_data , updating and retrievals,storage
P0PORT=4444
P1PORT=4454
P2PORT=4464
CAPORT=4400
PEERPEM=organizations/peerOrganizations/water_quality_data , updating and retrievals,storage/tlsca/tlsca.water_quality_data , updating and retrievals,storage-cert.pem
CAPEM=organizations/peerOrganizations/water_quality_data , updating and retrievals,storage/ca/ca.water_quality_data , updating and retrievals,storage-cert.pem

echo "$(json_ccp $ORG $PEER $P0PORT $P1PORT $P2PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/water_quality_data , updating and retrievals,storage/connection-water_quality_data , updating and retrievals,storage.json
echo "$(yaml_ccp $ORG $PEER $P0PORT $P1PORT $P2PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/water_quality_data , updating and retrievals,storage/connection-water_quality_data , updating and retrievals,storage.yaml
# save another copy of json connection profile in a different directory
echo "$(json_ccp $ORG $PEER $P0PORT $P1PORT $P2PORT $CAPORT $PEERPEM $CAPEM)" > network-config/network-config-water_quality_data , updating and retrievals,storage.json

## prepare connection profile for orggovernment
ORG=Government
PEER=government
P0PORT=5555
P1PORT=5565
P2PORT=5575
CAPORT=5500
PEERPEM=organizations/peerOrganizations/government/tlsca/tlsca.government-cert.pem
CAPEM=organizations/peerOrganizations/government/ca/ca.government-cert.pem

echo "$(json_ccp $ORG $PEER $P0PORT $P1PORT $P2PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/government/connection-government.json
echo "$(yaml_ccp $ORG $PEER $P0PORT $P1PORT $P2PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/government/connection-government.yaml
# save another copy of json connection profile in a different directory
echo "$(json_ccp $ORG $PEER $P0PORT $P1PORT $P2PORT $CAPORT $PEERPEM $CAPEM)" > network-config/network-config-government.json

## prepare connection profile for orgincentives_contrants
ORG=Incentives_contrants
PEER=incentives_contrants
P0PORT=6666
P1PORT=6676
P2PORT=6686
CAPORT=6600
PEERPEM=organizations/peerOrganizations/incentives_contrants/tlsca/tlsca.incentives_contrants-cert.pem
CAPEM=organizations/peerOrganizations/incentives_contrants/ca/ca.incentives_contrants-cert.pem

echo "$(json_ccp $ORG $PEER $P0PORT $P1PORT $P2PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/incentives_contrants/connection-incentives_contrants.json
echo "$(yaml_ccp $ORG $PEER $P0PORT $P1PORT $P2PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/incentives_contrants/connection-incentives_contrants.yaml
# save another copy of json connection profile in a different directory
echo "$(json_ccp $ORG $PEER $P0PORT $P1PORT $P2PORT $CAPORT $PEERPEM $CAPEM)" > network-config/network-config-incentives_contrants.json




