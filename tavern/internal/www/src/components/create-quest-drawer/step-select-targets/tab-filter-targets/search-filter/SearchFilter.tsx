import React from "react";
import Select from "react-select"

type SearchFilterParams = {
    beacons: Array<any>;
    setFilteredBeacons: (arg1: any) => void;
}
export const SearchFilter = (props: SearchFilterParams) => {
    const {beacons, setFilteredBeacons} = props;

    const options = [
        {
        label: "Service",
        options:[
            { value: "Relay", label: "Relay", type: "service" },
            { value: "Admin", label: "Admin", type: "service" },
            { value: "Web", label: "Web", type: "service" },
        ]
        },
        {
            label: "Group",
            options:[
                { value: "Team 1", label: "Team 1", type: "group" },
                { value: "Team 2", label: "Team 2", type: "group" },
                { value: "Team 3", label: "Team 3", type: "group" },
                { value: "Team 4", label: "Team 4", type: "group" },
                { value: "Team 5", label: "Team 5", type: "group" },
                { value: "Team 6", label: "Team 6", type: "group"}
            ]
        },
        {
            label: "Beacon",
            options:[
                { value: "15b9ec70-b3db-11ed-afa1-0242ac120002", label: "15b9ec70-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "15b9f04e-b3db-11ed-afa1-0242ac120002", label: "15b9f04e-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "15b9f99a-b3db-11ed-afa1-0242ac120002", label: "15b9f99a-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "15b9fd82-b3db-11ed-afa1-0242ac120002", label: "15b9fd82-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "15b9ffb2-b3db-11ed-afa1-0242ac120002", label: "15b9ffb2-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "25b9ffb2-b3db-11ed-afa1-0242ac120002", label: "25b9ffb2-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "35b9ffb2-b3db-11ed-afa1-0242ac120002", label: "35b9ffb2-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "45b9ffb2-b3db-11ed-afa1-0242ac120002", label: "45b9ffb2-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "55b9ffb2-b3db-11ed-afa1-0242ac120002", label: "55b9ffb2-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "65b9ffb2-b3db-11ed-afa1-0242ac120002", label: "65b9ffb2-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "75b9ffb2-b3db-11ed-afa1-0242ac120002", label: "75b9ffb2-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "85b9ffb2-b3db-11ed-afa1-0242ac120002", label: "85b9ffb2-b3db-11ed-afa1-0242ac120002", type: "beacon"},
                { value: "95b9ffb2-b3db-11ed-afa1-0242ac120002", label: "95b9ffb2-b3db-11ed-afa1-0242ac120002", type: "beacon"}

            ]
        }
    ];

    const handleChange = (selected: any) => {
        if(selected.length < 1 ){
            setFilteredBeacons(beacons);
        }
        else{
            const searchTypes = selected.reduce((accumulator:any, currentValue:any) => {
                if(currentValue.type === "beacon"){
                    accumulator.beacon.push(currentValue.value);
                }
                else if(currentValue.type === "service"){
                    accumulator.service.push(currentValue.value);
                }
                else if(currentValue.type === "group"){
                    accumulator.group.push(currentValue.value);
                }
                return accumulator;
            },
            {
                "beacon": [],
                "service": [],
                "group": []
            });

            const filtered = beacons.filter( (beacon) => {
                let group = (beacon?.tags).find( (obj : any) => {
                    return obj?.kind === "group"
                }) || null;

                let service = (beacon?.tags).find( (obj : any) => {
                    return obj?.kind === "service"
                }) || null;

                let match = true;

                if(searchTypes.beacon.length > 0){
                    // If a beacon filter is applied ignore other filters to just match the beacon
                    if(searchTypes.beacon.indexOf(beacon.id) > -1){
                        match = true;
                    }
                    else{
                        return false;
                    }
                }

                if(searchTypes.service.length > 0){
                    if(service && searchTypes.service.indexOf(service?.id) > -1){
                        match = true;
                    }
                    else{
                        return false;
                    }
                }

                if(searchTypes.group.length > 0){
                    if(group && searchTypes.group.indexOf(group?.id) > -1){
                        match = true;
                    }
                    else{
                        return false;
                    }
                }

                return match;
            });
            setFilteredBeacons(filtered);
        }
    };

    return (
        <Select
            isSearchable={true}
            isMulti
            options={options}
            onChange={handleChange}
        />
    );
};