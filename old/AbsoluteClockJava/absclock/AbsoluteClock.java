package absclock;

public class AbsoluteClock {
	private AbsoluteTime time;
	
	public AbsoluteTime getTime() {
		return time;
	}

	public void setTime(AbsoluteTime time) {
		this.time = time;
	}

	public static void main(String[] args) {
		new AbsoluteClock();
	}

	/**
	 * Constructor.
	 *
	 */
	public AbsoluteClock() {
		time = new AbsoluteTime();
		System.out.println(time);
	}
}