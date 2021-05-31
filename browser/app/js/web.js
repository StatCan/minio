/*
 * MinIO Object Storage (c) 2021 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import JSONrpc from './jsonrpc'
import { minioBrowserPrefix } from './constants.js'
import Moment from 'moment'
import storage from 'local-storage-fallback'
import i18n from './../i18n'
class Web {
  constructor(endpoint) {
    const namespace = 'web'
    this.JSONrpc = new JSONrpc({
      endpoint,
      namespace
    })
  }
  makeCall(method, options) {
    return this.JSONrpc.call(method, {
      params: options
    }, storage.getItem('token'))
      .catch(err => {
        if (err.status === 401) {
          storage.removeItem('token')
          location.reload()
          throw new Error(i18n.t('errReLogin'))
        }
        if (err.status)
          throw new Error(i18n.t('errReturnedError', { errorStatus: err.status }))
        throw new Error(i18n.t('errMinioUnreachable'))
      })
      .then(res => { //Rpc server handled the request
        let json = JSON.parse(res.text)
        let result = json.result
        let error = json.error
        if (error) { // This limits to only the message. 
          //If there is a code (numbered), show a generic error
          if(error.code){
            throw new Error(i18n.t('errGeneric'))
          }
          // If an error key was provided, find the error and translation
          if(error.data){
            let msg = this.createTranslatedError(error)
            throw new Error(msg)
          }
          throw new Error(i18n.t('errGeneric') + ' ' + error.message)
        }
        if (!Moment(result.uiVersion).isValid()) {
          throw new Error(i18n.t('errInvalidUIVersion'))
        }
        if (result.uiVersion !== currentUiVersion
          && currentUiVersion !== 'MINIO_UI_VERSION') {
          storage.setItem('newlyUpdated', true)
          location.reload()
        }
        return result
      })
  }
  createTranslatedError(error) {
    let key = error.data['key']
    let parameters = error.data['param']
    if(parameters){ 
      //There are parameter(s) for the message
      return i18n.t(key, parameters)
    }
    return i18n.t(key)
  }
  LoggedIn() {
    return !!storage.getItem('token')
  }
  Login(args) {
    return this.makeCall('Login', args)
      .then(res => {
        storage.setItem('token', `${res.token}`)
        return res
      })
  }
  Logout() {
    storage.removeItem('token')
  }
  GetToken() {
    return storage.getItem('token')
  }
  GetDiscoveryDoc() {
    return this.makeCall("GetDiscoveryDoc")
  }
  LoginSTS(args) {
    return this.makeCall('LoginSTS', args)
      .then(res => {
        storage.setItem('token', `${res.token}`)
        return res
      })
  }
  ServerInfo() {
    return this.makeCall('ServerInfo')
  }
  StorageInfo() {
    return this.makeCall('StorageInfo')
  }
  ListBuckets() {
    return this.makeCall('ListBuckets')
  }
  MakeBucket(args) {
    return this.makeCall('MakeBucket', args)
  }
  DeleteBucket(args) {
    return this.makeCall('DeleteBucket', args)
  }
  ListObjects(args) {
    return this.makeCall('ListObjects', args)
  }
  PresignedGet(args) {
    return this.makeCall('PresignedGet', args)
  }
  PutObjectURL(args) {
    return this.makeCall('PutObjectURL', args)
  }
  RemoveObject(args) {
    return this.makeCall('RemoveObject', args)
  }
  SetAuth(args) {
    return this.makeCall('SetAuth', args)
      .then(res => {
        storage.setItem('token', `${res.token}`)
        return res
      })
  }
  CreateURLToken() {
    return this.makeCall('CreateURLToken')
  }
  GetBucketPolicy(args) {
    return this.makeCall('GetBucketPolicy', args)
  }
  SetBucketPolicy(args) {
    return this.makeCall('SetBucketPolicy', args)
  }
  ListAllBucketPolicies(args) {
    return this.makeCall('ListAllBucketPolicies', args)
  }
}

const web = new Web(`${window.location.protocol}//${window.location.host}${minioBrowserPrefix}/webrpc`);

export default web;
