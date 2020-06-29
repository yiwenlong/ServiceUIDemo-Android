// IRemoteService.aidl
package com.yiwenlong.service;

// Declare any non-default types here with import statements
import com.yiwenlong.service.IRemoteServiceCallback;
import com.yiwenlong.service.IRemoteServiceLogCallback;

interface IRemoteService {

    void registerServiceCallback(IRemoteServiceCallback callback);

    void registerLogCallback(IRemoteServiceLogCallback callback);

    void bootServer(String host, int port);

    void stopServer();

    boolean isServerRunning();

    int getPid();
}
