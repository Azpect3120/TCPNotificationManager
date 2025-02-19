# TCPNotificationManager

<!--toc:start-->
- [TCPNotificationManager](#tcpnotificationmanager)
  - [The Problem](#the-problem)
  - [The Solution](#the-solution)
  - [Dependencies](#dependencies)
      - [Linux](#linux)
  - [Events](#events)
<!--toc:end-->

## The Problem

I have a server that I want to be able to notify me when it is done with a task. I don't want 
to have to poll the server to see if it is done. I want the server to notify me when it is done.
This server is a headless server, so I can't use a GUI to notify me. I want to be able to run
a command on the server to notify me directly. 

I also want all my notifications to be in one place. I don't want to have to check multiple
places for notifications. Furthremore, each system I use should be able to connect at the same
time and receive the same notifications at the same time. 


## The Solution
The server will host a TCP server on the internet that the clients (my local systems) can connect 
to and communicate with each other and the server using TCP with my own event system.

This application is purpose built for my own setup and my own systems. It is not meant to be used
by other people, however, if you would like to use it, you are welcome to. But, I will not include 
installation instructions for other systems.

If you have experience with Linux and GoLang, setting up this application should be easy.

## Dependencies

#### Linux

- notify-send

<!-- EVENTS_START -->
<!-- EVENTS_END -->
