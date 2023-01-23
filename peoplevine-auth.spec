%define debug_package %{nil}
%undefine _missing_build_ids_terminate_build

Name:    peoplevine-auth
Version: 0.0.1
Release: 1%{?dist}
Summary: PeopleVine auth proxy server
License: Private
Source0:  peoplevine-auth.tgz
BuildRequires: golang
BuildRequires: git
BuildRequires: systemd

%description
 PeopleVine auth proxy server

%prep
%setup -q -n peoplevine-auth

%build

export GOPATH=$(pwd) &&\
cd peoplevine-auth &&\
go build .

%install
install -d %{buildroot}%{_sbindir}
install -d %{buildroot}%{_unitdir}
install -d %{buildroot}%{_sysconfdir}
install -p -m 0755 peoplevine-auth/peoplevine-auth %{buildroot}%{_sbindir}
install -p -m 0644 systemd/peoplevine-auth.service %{buildroot}%{_unitdir}/peoplevine-auth.service

%post
%systemd_post peoplevine-auth.service

%preun
%systemd_preun peoplevine-auth.service

%postun
%systemd_postun_with_restart peoplevine-auth.service

%files
%doc README.md
%{_sbindir}/*
%{_unitdir}/*

%changelog
* Mon Jan 23 2023 Eduard Ponomarenko <gotik@mail.com> - 0.0.1-1
- initial version

