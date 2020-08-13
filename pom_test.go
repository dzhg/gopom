package gopom

import "testing"
import "github.com/stretchr/testify/assert"

func TestParsePomSimple(t *testing.T) {
	s := `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <parent>
    <groupId>org.sonatype.oss</groupId>
    <artifactId>oss-parent</artifactId>
    <version>7</version>
  </parent>

  <properties>
    <scala.version>2.10.0</scala.version>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
    <project.reporting.outputEncoding>UTF-8</project.reporting.outputEncoding>
    <pax.exam.version>2.6.0</pax.exam.version>
    <pax.url.version>1.5.0</pax.url.version>
  </properties>
	
  <groupId>org.scalatest</groupId>
  <artifactId>scalatest_2.10</artifactId>
  <version>2.0.M6-SNAP24</version>
  <packaging>jar</packaging>

  <name>ScalaTest</name>
  <description>
    ScalaTest is a free, open-source testing toolkit for Scala and Java
    programmers.
  </description>
  <url>http://www.scalatest.org</url>

  <licenses>
    <license>
      <name>the Apache License, ASL Version 2.0</name>
      <url>http://www.apache.org/licenses/LICENSE-2.0</url>
      <distribution>repo</distribution>
    </license>
  </licenses>

  <scm>
    <url>https://github.com/scalatest/scalatest</url>
    <connection>scm:git:git://github.com/scalatest/scalatest.git</connection>
    <developerConnection>
      scm:git:git@github.com:scalatest/scalatest.git
    </developerConnection>
  </scm>

  <developers>
    <developer>
      <id>bill</id>
      <name>Bill Venners</name>
      <email>bill@artima.com</email>
    </developer>
  </developers>

  <inceptionYear>2009</inceptionYear>

  <dependencies>

    <dependency>
      <groupId>org.scala-lang</groupId>
      <artifactId>scala-library</artifactId>
      <version>2.10.0</version>
    </dependency>

    <dependency>
      <groupId>org.scala-lang</groupId>
      <artifactId>scala-compiler</artifactId>
      <version>2.10.0</version>
      <optional>true</optional>
    </dependency>

    <dependency>
      <groupId>org.scalatest</groupId>
      <artifactId>test-interface</artifactId>
      <version>1.0-SNAP3</version>
      <optional>true</optional>
    </dependency>

    <dependency>
      <groupId>org.antlr</groupId>
      <artifactId>stringtemplate</artifactId>
      <version>3.2</version>
      <optional>true</optional>
    </dependency>

    <dependency>
      <groupId>org.scalacheck</groupId>
      <artifactId>scalacheck_2.10</artifactId>
      <version>1.10.0</version>
      <optional>true</optional>
    </dependency>

    <dependency>
      <groupId>org.easymock</groupId>
      <artifactId>easymockclassextension</artifactId>
      <version>3.1</version>
      <optional>true</optional>
    </dependency>

    <dependency>
      <groupId>org.jmock</groupId>
      <artifactId>jmock-legacy</artifactId>
      <version>2.5.1</version>
      <optional>true</optional>
    </dependency>

    <dependency>
      <groupId>org.mockito</groupId>
      <artifactId>mockito-all</artifactId>
      <version>1.9.0</version>
      <optional>true</optional>
    </dependency>

    <dependency>
      <groupId>org.testng</groupId>
      <artifactId>testng</artifactId>
      <version>6.3.1</version>
      <optional>true</optional>
    </dependency>
    
    <dependency>
      <groupId>com.google.inject</groupId>
      <artifactId>guice</artifactId>
      <version>3.0</version>
      <optional>true</optional>
    </dependency>

    <dependency>
      <groupId>junit</groupId>
      <artifactId>junit</artifactId>
      <version>4.10</version>
      <optional>true</optional>
    </dependency>
    
    <dependency>
      <groupId>org.seleniumhq.selenium</groupId>
      <artifactId>selenium-java</artifactId>
      <version>2.31.0</version>
      <optional>true</optional>
    </dependency>

    <dependency>
      <groupId>net.sourceforge.cobertura</groupId>
      <artifactId>cobertura</artifactId>
      <version>1.9.1</version>
      <scope>test</scope>
    </dependency>

    <dependency>
      <groupId>commons-io</groupId>
      <artifactId>commons-io</artifactId>
      <version>1.3.2</version>
      <scope>test</scope>
    </dependency>
    
    <dependency>
      <groupId>org.eclipse.jetty</groupId>
      <artifactId>jetty-server</artifactId>
      <version>8.1.8.v20121106</version>
      <scope>test</scope>
    </dependency>
    
    <dependency>
      <groupId>org.eclipse.jetty</groupId>
      <artifactId>jetty-webapp</artifactId>
      <version>8.1.8.v20121106</version>
      <scope>test</scope>
    </dependency>
    
    <dependency>
	  <groupId>asm</groupId>
	  <artifactId>asm</artifactId>
	  <version>3.3.1</version>
	  <optional>true</optional>
    </dependency>
    
    <dependency>
      <groupId>org.pegdown</groupId>
      <artifactId>pegdown</artifactId>
      <version>1.1.0</version>
      <optional>true</optional>
    </dependency>

    <dependency>
      <groupId>org.ops4j.pax.exam</groupId>
      <artifactId>pax-exam-container-native</artifactId>
      <version>${pax.exam.version}</version>
      <scope>test</scope>
    </dependency>

    <dependency>
      <groupId>org.ops4j.pax.exam</groupId>
      <artifactId>pax-exam-junit4</artifactId>
      <version>${pax.exam.version}</version>
      <scope>test</scope>
    </dependency>

    <dependency>
      <groupId>org.ops4j.pax.exam</groupId>
      <artifactId>pax-exam-link-mvn</artifactId>
      <version>${pax.exam.version}</version>
      <scope>test</scope>
    </dependency>

    <dependency>
      <groupId>org.ops4j.pax.url</groupId>
      <artifactId>pax-url-aether</artifactId>
      <version>${pax.url.version}</version>
      <scope>test</scope>
    </dependency>

    <dependency>
      <groupId>org.apache.felix</groupId>
      <artifactId>org.apache.felix.framework</artifactId>
      <version>4.0.3</version>
      <scope>test</scope>
    </dependency>

    <dependency>
      <groupId>org.eclipse.tycho</groupId>
      <artifactId>org.eclipse.osgi</artifactId>
      <version>3.7.0.v20110613</version>
      <scope>test</scope>
    </dependency>

    <dependency>
      <groupId>org.slf4j</groupId>
      <artifactId>slf4j-api</artifactId>
      <version>1.7.1</version>
      <scope>test</scope>
    </dependency>

    <dependency>
      <groupId>ch.qos.logback</groupId>
      <artifactId>logback-core</artifactId>
      <version>1.0.7</version>
      <scope>test</scope>
    </dependency>

    <dependency>
      <groupId>ch.qos.logback</groupId>
      <artifactId>logback-classic</artifactId>
      <version>1.0.7</version>
      <scope>test</scope>
    </dependency>

  </dependencies>

</project>`

	p, err := UnmarshalPOM([]byte(s))
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, "org.scalatest", p.GroupID)
	assert.Equal(t, "scalatest_2.10", p.ArtifactID)
	assert.Len(t, p.Dependencies, 27)
}

func TestParseDependency(t *testing.T) {
	s := `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
  <modelVersion>4.0.0</modelVersion>
  <dependencies>
    <dependency>
      <groupId>org.some.group</groupId>
      <artifactId>some_artifact</artifactId>
      <version>1.0.0</version>
      <optional>true</optional>
      <scope>test</scope>
    </dependency>
    <dependency>
      <groupId>org.some.group</groupId>
      <artifactId>other_artifact</artifactId>
      <version>1.0.1</version>
      <scope>compile</scope>
    </dependency>
  </dependencies>
</project>
`
	p, err := UnmarshalPOM([]byte(s))
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Len(t, p.Dependencies, 2)

	d := p.Dependencies[0]
	assert.Equal(t, "org.some.group", d.GroupID)
	assert.Equal(t, "some_artifact", d.ArtifactID)
	assert.Equal(t, "1.0.0", d.Version)
	assert.Equal(t, "test", d.Scope)
	assert.True(t, d.Optional)

	d = p.Dependencies[1]
	assert.Equal(t, "org.some.group", d.GroupID)
	assert.Equal(t, "other_artifact", d.ArtifactID)
	assert.Equal(t, "1.0.1", d.Version)
	assert.Equal(t, "compile", d.Scope)
	assert.False(t, d.Optional)
}

func TestParseDependencyManagement(t *testing.T) {
	s := `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
  <modelVersion>4.0.0</modelVersion>
  <dependencyManagement>
    <dependencies>
      <dependency>
        <groupId>org.some.group</groupId>
        <artifactId>some_artifact</artifactId>
        <version>1.0.0</version>
      </dependency>
    </dependencies>
  </dependencyManagement>
  <dependencies>
    <dependency>
      <groupId>org.some.group</groupId>
      <artifactId>some_artifact</artifactId>
    </dependency>
  </dependencies>
</project>
`
	p, err := UnmarshalPOM([]byte(s))
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Len(t, p.DependencyManagement.Dependencies, 1)
}

func TestParseProperties(t *testing.T) {
	s := `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
  <modelVersion>4.0.0</modelVersion>
  <properties>
    <test.property.a>a</test.property.a>
    <test.property.b>bb</test.property.b>
    <!-- this is a comment -->
    <test.property.c>ccc</test.property.c>
  </properties>
</project>
`
	p, err := UnmarshalPOM([]byte(s))
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Len(t, p.Properties.Entries, 3)

	keys := []string{"test.property.a", "test.property.b", "test.property.c"}
	values := []string{"a", "bb", "ccc"}
	for i := range p.Properties.Entries {
		assert.Equal(t, keys[i], p.Properties.Entries[i].Name)
		assert.Equal(t, values[i], p.Properties.Entries[i].Value)
	}
}

func TestParseExclusions(t *testing.T) {
	s := `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
  <modelVersion>4.0.0</modelVersion>
  <dependencies>
    <dependency>
      <groupId>abc.xyz</groupId>
      <artifactId>tuv</artifactId>
      <version>1.0.0</version>
      <exclusions>
        <exclusion>
          <groupId>abc.def.xxx</groupId>
          <artifactId>xxx</artifactId>
        </exclusion>
        <exclusion>
          <groupId>abc.def.yyy</groupId>
          <artifactId>yyy</artifactId>
        </exclusion>
      </exclusions>
    </dependency>
  </dependencies>
</project>
`
	p, err := UnmarshalPOM([]byte(s))
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Len(t, p.Dependencies, 1)
	assert.Len(t, p.Dependencies[0].Exclusions, 2)

	gids := []string{"abc.def.xxx", "abc.def.yyy"}
	aids := []string{"xxx", "yyy"}

	for i := range p.Dependencies[0].Exclusions {
		assert.Equal(t, gids[i], p.Dependencies[0].Exclusions[i].GroupID)
		assert.Equal(t, aids[i], p.Dependencies[0].Exclusions[i].ArtifactID)
	}
}
