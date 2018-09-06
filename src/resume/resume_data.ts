export default {
	profile: {
		name: {
			first: 'Jack',
			last: 'Sorrell'
		},
		title: 'Software Engineer',
		birthdate: '1993-12-22',
		residence: {
			city: 'Versailles',
			state: 'Kentucky',
			country: 'USA'
		},
		bio: 'I am an all-round developer who has yet to specialize. I have a strong understanding of the principles underlying all realms of computer science, rather than simply knowing all the intricacies of a single language. I prioritize robustness and correctness over implementation speed. I am an open-source enthusiast and see morality and privacy concerns as paramount.',
		image: '/images/myface-nobg.png'
	},
	degrees: [{
		name: 'Carnegie Mellon University',
		start_date: '2012-08',
		end_date: '2017-08',
		location: {
			city: 'Pittsburgh',
			state: 'Pennsylvania'
		},
		type: 'undergraduate',
		degree: 'Bachelor of Science',
		degree_short: 'BS',
		major: 'Computer Science',
		minors: [
			'Physics',
			'Mathematics'
		]
	}],
	courses: {
		'Computer Science': [{
			name: 'Compiler Design',
			shortName: 'Compiler Design',
			id: '15-411',
			description: 'This course covers the design and implementation of compiler and run-time systems for high-level languages, and examines the interaction between language design, compiler design, and run-time organization. Topics covered include syntactic and lexical analysis, handling of user-defined types and type-checking, context analysis, code generation and optimization, and memory management and run-time organization.'
		},
		{
			name: 'Operating System Design and Implementation',
			shortName: 'Operating Systems',
			id: '15-410',
			description: 'Operating System Design and Implementation is a rigorous hands-on introduction to the principles and practice of operating systems. The core experience is writing a small Unix-inspired OS kernel, in C with some x86 assembly language, which runs on a PC hardware simulator (and on actual PC hardware if you wish). Work is done in two-person teams, and \'team programming\' skills (source control, modularity, documentation) are emphasized. The size and scope of the programming assignments typically result in students significantly developing their design, implementation, and debugging abilities. Core concepts include the process model, virtual memory, threads, synchronization, and deadlock; the course also surveys higher-level OS topics including file systems, interprocess communication, networking, and security.'
		},
		{
			name: 'Algorithm Design and Analysis',
			shortName: 'Algorithm Design',
			id: '15-451',
			description: 'This course is about the design and analysis of algorithms. We study specific algorithms for a variety of problems, as well as general design and analysis techniques. Specific topics include searching, sorting, algorithms for graph problems, efficient data structures, lower bounds and NP-completeness. A variety of other topics may be covered at the discretion of the instructor. These include parallel algorithms, randomized algorithms, geometric algorithms, low level techniques for efficient programming, cryptography, and cryptographic protocols.'
		},
		{
			name: 'Foundations of Programming Languages',
			shortName: 'Programming Languages',
			id: '15-312',
			description: 'This course discusses in depth many of the concepts underlying the design, definition, implementation, and use of modern programming languages. Formal approaches to defining the syntax and semantics are used to describe the fundamental concepts underlying programming languages. A variety of programming paradigms are covered such as imperative, functional, logic, and concurrent programming. In addition to the formal studies, experience with programming in the languages is used to illustrate how different design goals can lead to radically different languages and models of computation.'
		},
		{
			name: 'Introduction to Computer Systems',
			shortName: 'Computer Systems',
			id: '15-213',
			description: 'This course provides a programmer\'s view of how computer systems execute programs, store information, and communicate. It enables students to become more effective programmers, especially in dealing with issues of performance, portability and robustness. It also serves as a foundation for courses on compilers, networks, operating systems, and computer architecture, where a deeper understanding of systems-level issues is required. Topics covered include: machine-level code and its generation by optimizing compilers, performance evaluation and optimization, computer arithmetic, memory organization and management, networking technology and protocols, and supporting concurrent computation.'
		},
		{
			name: 'Parallel and Sequential Data Structures and Algorithms',
			shortName: 'Data Structures and Algorithms',
			id: '15-210',
			description: 'Teaches students about how to design, analyze, and program algorithms and data structures. The course emphasizes parallel algorithms and analysis, and how sequential algorithms can be considered a special case. The course goes into more theoretical content on algorithm analysis than 15-122 and 15-150 while still including a significant programming component and covering a variety of practical applications such as problems in data analysis, graphics, text processing, and the computational sciences.'
		},
		{
			name: 'Principles of Functional Programming',
			shortName: 'Functional Programming',
			id: '15-150',
			description: 'One major theme is the interplay between inductive types, which are built up incrementally; recursive functions, which compute over inductive types by decomposition; and proof by structural induction, which is used to prove the correctness and time complexity of a recursive function. Another major theme is the role of types in structuring large programs into separate modules, and the integration of imperative programming through the introduction of data types whose values may be altered during computation.'
		},
		{
			name: 'Database Management Systems',
			shortName: 'DB Management Systems',
			id: 'CS1555',
			description: 'There are two principle objectives for this course. First, to introduce the fundamental concepts necessary for the design and use of a database.  Second, to provide practical experience in applying these concepts using commercial database management systems.'
		}],
		Mathematics: [{
			name: 'Number Theory',
			shortName: 'Number Theory',
			id: '21-441',
			description: 'Number theory deals with the integers, the most basic structures of mathematics. It is one of the most ancient, beautiful, and well-studied branches of mathematics, and has recently found surprising new applications in communications and cryptography. Course contents: Structure of the integers, greatest common divisiors, prime factorization. Modular arithmetic, Fermat\'s Theorem, Chinese Remainder Theorem. Number theoretic functions, e.g. Euler\'s function, Mobius functions, and identities. Diophantine equations, Pell\'s Equation, continued fractions. Modular polynomial equations, quadratic reciprocity.'
		},
		{
			name: 'Basic Logic',
			shortName: 'Logic',
			id: '21-300',
			description: 'Propositional and predicate logic: Syntax, proof theory and semantics up to completeness theorem, Lowenheim Skolem theorems, and applications of the compactness theorem.'
		},
		{
			name: 'Combinatorics',
			shortName: 'Combinatorics',
			id: '21-301',
			description: 'A major part of the course concentrates on algebraic methods, which are relevant in the study of error correcting codes, and other areas. Topics covered in depth include permutations and combinations, generating functions, recurrence relations, the principle of inclusion and exclusion, and the Fibonacci sequence and the harmonic series.'
		},
		{
			name: 'Matrix Theory',
			shortName: 'Matrix Theory',
			id: '21-242',
			description: 'An honors version of 21-241 (Matrix Algebra and Linear Transformations) for students of greater aptitude and motivation. More emphasis will be placed on writing proofs. Topics to be covered: complex numbers, real and complex vectors and matrices, rowspace and columnspace of a matrix, rank and nullity, solving linear systems by row reduction of a matrix, inverse matrices and determinants, change of basis, linear transformations, inner product of vectors, orthonormal bases and the Gram-Schmidt process, eigenvectors and eigenvalues, diagonalization of a matrix, symmetric and orthogonal matrices, hermitian and unitary matrices, quadratic forms.'
		},
		{
			name: 'Algebraic Structures',
			shortName: 'Algebraic Structures',
			id: '21-373',
			description: 'Groups: Homomorphisms. Subgroups, cosets, Lagrange\'s theorem. Conjugation. Normal subgroups, quotient groups, first isomorphism theorem. Group actions, Cauchy\'s Theorem. Dihedral and alternating groups. The second and third isomorphism theorems. Rings: Subrings, ideals, quotient rings, first isomorphism theorem. Polynomial rings. Prime and maximal ideals, prime and irreducible elements. PIDs and UFDs. Noetherian domains. Gauss\' lemma. Eisenstein criterion. Fields: Field of fractions of an integral domain. Finite fields. Applications to coding theory, cryptography, number theory.'
		},
		{
			name: 'Probability',
			shortName: 'Probability',
			id: '21-325',
			description: 'This course focuses on the understanding of basic concepts in probability theory and illustrates how these concepts can be applied to develop and analyze a variety of models arising in computational biology, finance, engineering and computer science. The firm grounding in the fundamentals is aimed at providing students the flexibility to build and analyze models from diverse applications as well as preparing the interested student for advanced work in these areas.'
		},
		{
			name: 'Calculus in Three Dimensions',
			shortName: '3D Calculus',
			id: '21-259',
			description: 'Vectors, lines, planes, quadratic surfaces, polar, cylindrical and spherical coordinates, partial derivatives, directional derivatives, gradient, divergence, curl, chain rule, maximum-minimum problems, multiple integrals, parametric surfaces and curves, line integrals, surface integrals, Green-Gauss theorems.'
		}],
		Physics: [{
			name: 'Physical Mechanics 1',
			shortName: 'Physical Mechanics',
			id: '33-331',
			description: 'Fundamental concepts of classical mechanics. Conservation laws, momentum, energy, angular momentum, Lagrange\'s and Hamilton\'s equations, motion under a central force, scattering, cross section, and systems of particles.'
		},
		{
			name: 'Quantum Physics',
			shortName: 'Quantum Physics',
			id: '33-234',
			description: 'An introduction to the fundamental principles and applications of quantum physics. A brief review of the experimental basis for quantization motivates the development of the Schrodinger wave equation. Several unbound and bound problems are treated in one dimension. The properties of angular momentum are developed and applied to central potentials in three dimensions. The one electron atom is then treated. Properties of collections of indistinguishable particles are developed allowing an understanding of the structure of the Periodic Table of elements.'
		},
		{
			name: 'Mathematical Methods of Physics',
			shortName: 'Physics Mathematics',
			id: '33-232',
			description: 'This course introduces, in the context of physical systems, a variety of mathematical tools and techniques that will be needed for later courses in the physics curriculum. Topics will include, linear algebra, vector calculus with physical application, Fourier series and integrals, partial differential equations and boundary value problems.'
		},
		{
			name: 'Intermediate Electricity and Magnetism 1',
			shortName: 'Electricity and Magnetism',
			id: '33-338',
			description: 'This course includes the basic concepts of electro- and magnetostatics. In electrostatics, topics include the electric field and potential for typical configurations, work and energy considerations, the method of images and solutions of Laplace\'s Equation, multipole expansions, and electrostatics in the presence of matter. In magnetostatics, the magnetic field and vector potential, magnetostatics in the presence of matter, properties of dia-, para- and ferromagnetic materials are developed.'
		},
		{
			name: 'Thermal Physics 1',
			shortName: 'Thermodynamics',
			id: '33-341',
			description: 'This course begins with a more systematic development of formal probability theory, with emphasis on generating functions, probability density functions and asymptotic approximations. Examples are taken from games of chance, geometric probabilities and radioactive decay. The connections between the ensembles of statistical mechanics (microcanonical, canonical and grand canonical) with the various thermodynamic potentials is developed for single component and multicomponent systems. Fermi-Dirac and Bose-Einstein statistics are reviewed. These principles are then applied to applications such as electronic specific heats, Einstein condensation, chemical reactions, phase transformations, mean field theories, binary phase diagrams, paramagnetism, ferromagnetism, defects, semiconductors and fluctuation phenomena.'
		},
		{
			name: 'Electronics 1',
			shortName: 'Electronics',
			id: '33-228',
			description: 'An introductory laboratory and lecture course with emphasis on elementary circuit analysis, design, and testing. We start by introducing basic circuit elements and study the responses of combinations to DC and AC excitations. We then take up transistors and learn about biasing and the behavior of amplifier circuits. The many uses of operational amplifiers are examined and analyzed; general features of feedback systems are introduced in this context. Complex functions are used to analyze all of the above linear systems. Finally, we examine and build some simple digital integrated circuits.'
		},
		{
			name: 'Introduction to Computational Physics',
			shortName: 'Computational Physics',
			id: '33-241',
			description: 'This course is designed to introduce the student to the computational tools needed for study and research in science, engineering and other disciplines. Topics covered include: types of computers and computing systems; programming languages and environments, applications and application areas with their appropriate algorithms; issues of correctness.'
		},
		{
			name: 'Physical Analysis',
			shortName: 'Physical Analysis',
			id: '33-231',
			description: 'This course aims to develop analytical skills and mathematical modeling skills across a broad spectrum of physical phenomena, stressing analogies in behavior of a wide variety of systems. Specific topics include dimensional analysis and scaling in physical phenomena, exponential growth and decay, the harmonic oscillator with damping and driving forces, linear approximations of nonlinear systems, coupled oscillators, and wave motion. Necessary mathematical techniques, including differential equations, complex exponential functions, matrix algebra, and elementary Fourier series, are introduced as needed.'
		}]
	},
	experiences: [{
		title: 'Software Engineering Intern',
		company: 'Applied Predictive Technologies',
		location: {
			city: 'Arlington',
			state: 'Virginia'
		},
		start_date: '2015-05',
		end_date: '2015-08',
		description: 'Applied Predictive Technologies (APT) is a SAAS company which seeks to allow business to systematically test changes on a subset of its businesses before deciding whether to implement everywhere. In this internship, I used the Agile approach to software development and worked mostly using C# and Microsoft SQL Server. I joined two teams: the Infrastructure team and an Analytics team. As part of the Infrastructure team, I lead the effort to implement a job scheduling platform. I designed and built every part of the system from the backend scheduling to the API to the client library and web interface. As part of the Analytics team, I implemented an improved calculation that was rolled out into APT\'s main product.'
	}],
	projects: [{
		name: 'www.jacksorrell.com',
		start_date: '2017-12',
		end_date: '2018-02',
		description: 'Believing I needed a formal online presence and more experience in current web technologies, I decided to build the website you are currently visiting. In the process, I taught myself Bootstrap, Pug, Sass, Express as well as brushed up on Javascript/NodeJS, HTML, CSS and Nginx. Source is available on GitHub.'
	},
	{
		name: 'C0 Language Compiler',
		start_date: '2015-08',
		end_date: '2015-12',
		description: 'As part of my Compiler Design course at Carnegie Mellon University, I developed a compiler for a C-like programming language taught in early computer science courses. I designed and implemented every part of the compiler from lexing to assembly generation. I used SSA form for the intermediate representation and implemented many optimizations including constant propogation and folding, dead-code elimination, tail-call optimization, function-inlining, and register coalescing. Source is available upon request.'
	},
	{
		name: 'OS Kernel',
		start_date: '2015-01',
		end_date: '2015-05',
		description: 'As part of my Operating System Design and Implementation course at Carnegie Mellon University, I worked with a partner to develop an x86 kernel. We implemented a thread-library for use by user processes, as well as virtual memory, context switching, and a disk driver. Source is available upon request.'
	},
	{
		name: 'Tetris Friends AI',
		start_date: '2014-04',
		end_date: '2014-09',
		description: 'Begun in a hackathon and completed individually, I designed an AI to play the game Tetris online at www.tetrisfriends.com. I used pcap to intercept TCP packets sent from the server containing a seed for the a random number generator that determined which pieces would be given in which order. By decompiling the local Java client, I reverse engineered the random number generator. I designed the AI to then determine possible future board states in a few rounds and use a heuristic to choose the optimal one. Moves for generating this board state were then converted into a series of virtual keypresses that were passed to the client program. Source is available on GitHub.'
	}],
	skills: {
		Platforms: {
			Linux: {
				'level': 5
			},
			Web: {
				'level': 4
			},
			Windows: {
				'level': 3
			},
			Android: {
				'level': 3
			}
		},
		'Development Concepts': {
			'Object Oriented Design': {
				'level': 5
			},
			'Functional Programming': {
				'level': 5
			},
			Recursion: {
				'level': 5
			},
			Algorithms: {
				'level': 5
			},
			Concurrency: {
				'level': 4
			},
			'Dynamic Programming': {
				'level': 3
			}
		},
		Languages: {
			C: {
				'level': 5
			},
			'C++': {
				'level': 4
			},
			'C#': {
				'level': 4
			},
			Javascript: {
				'level': 4
			},
			Haskell: {
				'level': 4
			},
			Python: {
				'level': 4
			},
			Java: {
				'level': 4
			},
			Go: {
				'level': 4
			}
		},
		Tools: {
			'Linux Command Line': {
				'level': 5
			},
			Atom: {
				level: 4
			},
			Vim: {
				'level': 4
			},
			Git: {
				'level': 4
			},
			Latex: {
				'level': 4
			},
			'Visual Studio': {
				'level': 3
			},
			'Intellij IDEA': {
				'level': 3
			}
		}
	},
	links: {
		github: {
			icon: '/images/octocat.svg',
			href: 'https://github.com/jsorrell'
		},
		twitter: {
			icon: '/images/Twitter_Social_Icon_Circle_Color.svg',
			href: 'https://twitter.com/jsorrell414'
		},
		keybase: {
			icon: '/images/keybase_logo_official.svg',
			href: 'https://keybase.io/jsorrell'
		}
	}
};
