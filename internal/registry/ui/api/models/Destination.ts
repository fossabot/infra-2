/* tslint:disable */
/* eslint-disable */
/**
 * Infra API
 * Infra REST API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    DestinationKubernetes,
    DestinationKubernetesFromJSON,
    DestinationKubernetesFromJSONTyped,
    DestinationKubernetesToJSON,
} from './';

/**
 * 
 * @export
 * @interface Destination
 */
export interface Destination {
    /**
     * 
     * @type {string}
     * @memberof Destination
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof Destination
     */
    name: string;
    /**
     * 
     * @type {number}
     * @memberof Destination
     */
    created: number;
    /**
     * 
     * @type {number}
     * @memberof Destination
     */
    updated: number;
    /**
     * 
     * @type {Array<string>}
     * @memberof Destination
     */
    labels: Array<string>;
    /**
     * 
     * @type {DestinationKubernetes}
     * @memberof Destination
     */
    kubernetes?: DestinationKubernetes;
}

export function DestinationFromJSON(json: any): Destination {
    return DestinationFromJSONTyped(json, false);
}

export function DestinationFromJSONTyped(json: any, ignoreDiscriminator: boolean): Destination {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'name': json['name'],
        'created': json['created'],
        'updated': json['updated'],
        'labels': json['labels'],
        'kubernetes': !exists(json, 'kubernetes') ? undefined : DestinationKubernetesFromJSON(json['kubernetes']),
    };
}

export function DestinationToJSON(value?: Destination | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'created': value.created,
        'updated': value.updated,
        'labels': value.labels,
        'kubernetes': DestinationKubernetesToJSON(value.kubernetes),
    };
}


