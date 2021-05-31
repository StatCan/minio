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

import React from "react"
import { connect } from "react-redux"
import { Modal, ModalHeader, ModalBody } from "react-bootstrap"
import CopyToClipboard from "react-copy-to-clipboard"
import web from "../web"
import * as objectsActions from "./actions"
import * as alertActions from "../alert/actions"
import {
  SHARE_OBJECT_EXPIRY_DAYS,
  SHARE_OBJECT_EXPIRY_HOURS,
  SHARE_OBJECT_EXPIRY_MINUTES
} from "../constants"
import QRCode from "react-qr-code";
import { withTranslation } from "react-i18next"

export class ShareObjectModal extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      expiry: {
        days: SHARE_OBJECT_EXPIRY_DAYS,
        hours: SHARE_OBJECT_EXPIRY_HOURS,
        minutes: SHARE_OBJECT_EXPIRY_MINUTES
      }
    }
    this.expiryRange = {
      days: { min: 0, max: 7 },
      hours: { min: 0, max: 23 },
      minutes: { min: 0, max: 59 }
    }
  }
  updateExpireValue(param, inc) {
    let expiry = Object.assign({}, this.state.expiry)

    // Not allowing any increments if days is already max
    if (expiry.days == this.expiryRange["days"].max && inc > 0) {
      return
    }

    const { min, max } = this.expiryRange[param]
    expiry[param] = expiry[param] + inc
    if (expiry[param] < min || expiry[param] > max) {
      return
    }

    if (expiry.days == this.expiryRange["days"].max) {
      expiry.hours = 0
      expiry.minutes = 0
    } else if (expiry.days + expiry.hours + expiry.minutes == 0) {
      expiry.days = this.expiryRange["days"].max
    }

    this.setState({
      expiry
    })

    const { object, shareObject } = this.props
    shareObject(object.name, expiry.days, expiry.hours, expiry.minutes)
  }
  onUrlCopied() {
    const { showCopyAlert, hideShareObject, t } = this.props
    showCopyAlert(t('msgLinkCopied'))
    hideShareObject()
  }
  render() {
    const { shareObjectDetails, hideShareObject, t } = this.props
    const url = `${window.location.protocol}//${shareObjectDetails.url}`
    return (
      <Modal
        show={true}
        animation={false}
        onHide={hideShareObject}
        bsSize="small"
      >
        <ModalHeader>{t('shareObject')}</ModalHeader>
        <ModalBody>
          <div className="input-group copy-text">
            <QRCode value={url} size={128}/>
            <label>{t('shareableLink')}</label>
            <input
              type="text"
              ref={node => (this.copyTextInput = node)}
              readOnly="readOnly"
              value={url}
              onClick={() => this.copyTextInput.select()}
            />
          </div>
          {shareObjectDetails.showExpiryDate && (
            <div
            className="input-group"
            style={{ display: web.LoggedIn() ? "block" : "none" }}
          >
            <label>{t('expire')}</label>
            <div className="set-expire">
              <div className="set-expire-item">
                <i
                  id="increase-days"
                  className="set-expire-increase"
                  onClick={() => this.updateExpireValue("days", 1)}
                />
                <div className="set-expire-title">{t('days')}</div>
                <div className="set-expire-value">
                  <input
                    ref="expireDays"
                    type="number"
                    min={0}
                    max={7}
                    value={this.state.expiry.days}
                    readOnly="readOnly"
                  />
                </div>
                <i
                  id="decrease-days"
                  className="set-expire-decrease"
                  onClick={() => this.updateExpireValue("days", -1)}
                />
              </div>
              <div className="set-expire-item">
                <i
                  id="increase-hours"
                  className="set-expire-increase"
                  onClick={() => this.updateExpireValue("hours", 1)}
                />
                <div className="set-expire-title">{t('hours')}</div>
                <div className="set-expire-value">
                  <input
                    ref="expireHours"
                    type="number"
                    min={0}
                    max={23}
                    value={this.state.expiry.hours}
                    readOnly="readOnly"
                  />
                </div>
                <i
                  className="set-expire-decrease"
                  id="decrease-hours"
                  onClick={() => this.updateExpireValue("hours", -1)}
                />
              </div>
              <div className="set-expire-item">
                <i
                  id="increase-minutes"
                  className="set-expire-increase"
                  onClick={() => this.updateExpireValue("minutes", 1)}
                />
                <div className="set-expire-title">{t('minutes')}</div>
                <div className="set-expire-value">
                  <input
                    ref="expireMins"
                    type="number"
                    min={0}
                    max={59}
                    value={this.state.expiry.minutes}
                    readOnly="readOnly"
                  />
                </div>
                <i
                  id="decrease-minutes"
                  className="set-expire-decrease"
                  onClick={() => this.updateExpireValue("minutes", -1)}
                />
              </div>
            </div>
          </div>
          )}
        </ModalBody>
        <div className="modal-footer">
          <CopyToClipboard
            text={url}
            onCopy={this.onUrlCopied.bind(this)}
          >
            <button className="btn btn-success">{t('copyLink')}</button>
          </CopyToClipboard>
          <button className="btn btn-link" onClick={hideShareObject}>
            {t('common:cancel')}
          </button>
        </div>
      </Modal>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    object: ownProps.object,
    shareObjectDetails: state.objects.shareObject
  }
}

const mapDispatchToProps = dispatch => {
  return {
    shareObject: (object, days, hours, minutes) =>
      dispatch(objectsActions.shareObject(object, days, hours, minutes)),
    hideShareObject: () => dispatch(objectsActions.hideShareObject()),
    showCopyAlert: message =>
      dispatch(alertActions.set({ type: "success", message: message }))
  }
}

export default withTranslation('shareObjectModal')(connect(mapStateToProps, mapDispatchToProps)(ShareObjectModal))
