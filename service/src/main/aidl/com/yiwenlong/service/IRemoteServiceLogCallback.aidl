// IRemoteServiceLogCallback.aidl
package com.yiwenlong.service;

// Declare any non-default types here with import statements

interface IRemoteServiceLogCallback {
    void onReceiveLog(String message);
}
