# Nexus

### **Introduction**

----

个人成长项目，用于完善、连接和熟悉个人技术栈。

----

## To start using Nexus



## To start developing Nexus

The [community repository] hosts all information about
building Nexus from source, how to contribute code
and documentation, who to contact about what, etc.

If you want to build Nexus right away there are two options:

##### You have a working [Go environment].

```
mkdir -p $GOPATH/src/k8s.io
cd $GOPATH/src/k8s.io
git clone https://github.com/Nexus/Nexus
cd Nexus
make
```

##### You have a working [Docker environment].

```
git clone https://github.com/Nexus/Nexus
cd Nexus
make quick-release
```

For the full story, head over to the [developer's documentation].

## Support

If you need support, start with the [troubleshooting guide],
and work your way through the process that we've outlined.

That said, if you have questions, reach out to us
[one way or another][communication].