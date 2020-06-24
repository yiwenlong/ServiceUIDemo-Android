package com.yiwenlong.service;

import android.app.Notification;
import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.Service;
import android.content.Context;
import android.content.Intent;
import android.os.IBinder;
import android.util.Log;

import androidx.annotation.Nullable;

public class ServerManagerService extends Service {

    private static final String CHANNEL_ID = "channel id";
    private static final String CHANNEL_NAME = "channel name";

    @Override
    public void onCreate() {
        Log.i("ServerManagerService", "onCreate");
        super.onCreate();
        NotificationChannel channel = new NotificationChannel(CHANNEL_ID,CHANNEL_NAME,
                NotificationManager.IMPORTANCE_HIGH);

        NotificationManager manager = (NotificationManager) getSystemService(Context.NOTIFICATION_SERVICE);
        manager.createNotificationChannel(channel);

        Notification notification = new Notification.Builder(getApplicationContext(),CHANNEL_ID).build();
        startForeground(1, notification);
    }

    @Nullable
    @Override
    public IBinder onBind(Intent intent) {
        return null;
    }

    @Override
    public int onStartCommand(Intent intent, int flags, int startId) {
        bridge.Bridge.bootServer("0.0.0.0", 8080);
        return super.onStartCommand(intent, flags, startId);
    }

    @Override
    public void onDestroy() {
        bridge.Bridge.stopServer(8080);
        super.onDestroy();
    }
}
