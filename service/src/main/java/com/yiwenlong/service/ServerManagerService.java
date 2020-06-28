package com.yiwenlong.service;

import android.app.Notification;
import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.Service;
import android.content.Intent;
import android.os.IBinder;
import android.util.Log;

import androidx.annotation.Nullable;
import androidx.core.app.NotificationCompat;
import androidx.core.app.NotificationManagerCompat;

import bridge.MyServerListener;

public class ServerManagerService extends Service implements MyServerListener {

    private static final String CHANNEL_ID = "channel id";

    @Override
    public void onCreate() {
        Log.i("ServerManagerService", "onCreate");
        super.onCreate();

        NotificationManagerCompat.from(this)
                .createNotificationChannel(
                        new NotificationChannel(
                                CHANNEL_ID,
                                "a demo channel",
                                NotificationManager.IMPORTANCE_DEFAULT
                        )
                );

        Notification notification = new NotificationCompat.Builder(getApplicationContext(),CHANNEL_ID)
                .setContentText("这是一个 Android 下运行 Golang server 的 App")
                .setContentTitle("ServiceUI Shell")
                .setSmallIcon(R.mipmap.ic_launcher_round)
                .setGroup(CHANNEL_ID)
                .build();
        startForeground(256, notification);
    }

    @Nullable
    @Override
    public IBinder onBind(Intent intent) {
        return null;
    }

    @Override
    public int onStartCommand(Intent intent, int flags, int startId) {
        boolean isStart = intent.getBooleanExtra("start", true);
        if (isStart) {
            bridge.Bridge.initServer("0.0.0.0", 8080);
            bridge.Bridge.registerLogHandler(s -> Log.i("MyDemoGolangServer", s));
            bridge.Bridge.registerServerListener(this);
            bridge.Bridge.bootServer();
        } else {
            bridge.Bridge.stopServer();
        }
        return super.onStartCommand(intent, flags, startId);
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
    }

    @Override
    public void onServerError(String msg) {
        Log.e("MyServerBirdge", "onServerError -> " + msg);
    }

    @Override
    public void onServerStart() {
        Log.e("MyServerBirdge", "onServerStart");
    }

    @Override
    public void onServerStop() {
        Log.e("MyServerBirdge", "onServerStop");
        stopSelf();
    }
}
